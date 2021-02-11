package chopsticks

import (
	"fmt"
	. "github.com/logrusorgru/aurora/v3"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// loadConfig retrieves a file path and return its content
func loadConfig(file string) (string, error) {
	config, err := ioutil.ReadFile(file)
	// TODO: more err handle
	if err != nil {
		panic(err)
	}

	return string(config), nil
}

// getConfig retrieve a string and return its value in
func getConfig(name string) string {
	return gjson.Get(scoopConfig, name).String()
}

func setConfig(name, value string) (string, error) {
	config, err := sjson.Set(scoopConfig, name, value)
	if err != nil {
		panic(err)
	}

	return config, nil
}

func Error(a string) {
	msg := Sprintf(Red("ERROR %v"), Red(a))

	fmt.Println(msg)
}

func Warn(a string) {
	msg := Sprintf(Yellow("WARN %v"), Yellow(a))

	fmt.Println(msg)
}

func Info(a string) {
	msg := Sprintf(Gray(12-1, "INFO %v"), Gray(12-1, a))

	fmt.Println(msg)
}

func Success(a string) {
	msg := Sprintf(Green("%v"), Green(a))

	fmt.Println(msg)
}

func HumanReadableFileSize(length int64) string {
	// 1 KB = 1024 B
	const unit = 1024
	if length < unit {
		return fmt.Sprintf("%d B", length)
	}
	div, exp := int64(unit), 0
	for n := length / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(length)/float64(div), "KMGTPE"[exp])
}

// dirs
func baseDir(isGlobal bool) string {
	if isGlobal {
		return GlobalDir
	}
	return ScoopDir
}

func appsDir(isGlobal bool) string {
	return filepath.Join(baseDir(isGlobal), "apps")
}

func shimDir(isGlobal bool) string {
	return filepath.Join(baseDir(isGlobal), "shims")
}

func appDir(app string, isGlobal bool) string {
	return filepath.Join(appsDir(isGlobal), app)
}

func versionDir(app, version string, isGlobal bool) string {
	return filepath.Join(appDir(app, isGlobal), version)
}

func persistDir(app string, isGlobal bool) string {
	return filepath.Join(baseDir(isGlobal), "persist", app)
}

func userManifestDir() string {
	return filepath.Join(ScoopDir, "workspace")
}

func userManifest(app string) string {
	appJson := app + ".json"
	return filepath.Join(userManifestDir(), appJson)
}

func cachePath(app, version, url string) string {
	re := regexp.MustCompile(`[^\w.\-]+`)
	urlReplaced := re.ReplaceAllString(url, "_")
	filename := fmt.Sprintf("%v#%v#%v", app, version, urlReplaced)
	return filepath.Join(CacheDir, filename)
}

func SanitaryPath(path string) string {
	re := regexp.MustCompile(`[/\\?:*<>|]`)

	s := re.ReplaceAllString(path, "")

	return s
}

func isAppInstalled(app string, isGlobal bool) (bool, error) {
	// Dependencies of the format "bucket/dependency" install in a directory of form
	// "dependency". So we need to extract the bucket from the name and only give the app
	// name to is_directory
	s := strings.Split(app, "/")
	appName := s[len(s)-1]
	dir := appDir(appName, isGlobal)
	fi, err := os.Stat(dir)
	if err != nil {
		return false, err
	}
	return fi.IsDir(), nil
}

func GetInstalledApps(isGlobal bool) ([]string, error) {
	dir := appsDir(isGlobal)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return []string{}, nil
	}

	files, err := ioutil.ReadDir(dir)
	var apps []string
	if err != nil {
		return []string{}, err
	}
	for _, file := range files {
		if file.IsDir() {
			apps = append(apps, file.Name())
		}
	}

	return apps, nil
}

func GetAppFilePath(app, file string) string {
	// normal path to file
	path := filepath.Join(versionDir(app, "current", false), file)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return path
	}

	// global path to file
	path = filepath.Join(versionDir(app, "current", true), file)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return path
	}

	// not found
	return ""
}

func isCommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// available helpers: 7zip, Lessmsi, Innounp, Dark, Aria2
// what I wanna is a Enum
func helperToExe(helper string) string {
	helpers := map[string]string{
		"7zip":    "7z",
		"lessmsi": "lessmsi",
		"innounp": "innounp",
		"dark":    "dark",
		"aria2":   "aria2c",
	}

	return helpers[strings.ToLower(helper)]
}

// available helpers: 7zip, Lessmsi, Innounp, Dark, Aria2
func GetHelperPath(helper string) string {
	helperPath, err := exec.LookPath(helperToExe(helper))
	if err != nil {
		return ""
	}

	return helperPath
}

func isHelperInstalled(helper string) bool {
	path := GetHelperPath(helper)
	return !(path == "")
}

func isAria2Enabled() bool {
	return isHelperInstalled("Aria2") && gjson.Get(scoopConfig, "aria2-enabled").Bool()
}

type AppStatus struct {
	Installed     bool
	Version       string
	LatestVersion string
	Failed        bool
	Hold          bool
	Removed       bool
	Outdated      bool
	MissingDeps   []string
}

func GetAppStatus(app string, isGlobal bool) (AppStatus, error) {
	var s AppStatus

	installed, err := isAppInstalled(app, isGlobal)
	if err != nil {
		return AppStatus{}, err
	}
	s.Installed = installed

	version, err := CurrentVersion(app, isGlobal)
	if err != nil {
		return AppStatus{}, err
	}
	s.Version = version
	s.LatestVersion = s.Version

	i, err := GetInstallInfo(app, s.Version, isGlobal)
	if err != nil {
		return AppStatus{}, err
	}

	s.Hold = i.Hold

	var mf Manifest

	if i.Bucket != "" {
		mf, err = GetManifest(app, i.Bucket)
		if err != nil {
			return AppStatus{}, err
		}
	} else {
		mf = GetUrlManifest(i.URL)
	}

	if mf.Version != "" {
		s.LatestVersion = mf.Version
	}

	if s.Version != "" && s.LatestVersion != "" {
		s.Outdated = compareVersions(s.LatestVersion, s.Version) > 0
	}

	return s, nil
}

// getFileName Returns the name and extension parts of the given path
// got from dotnet
// https://github.com/microsoft/referencesource/blob/5697c29004a34d80acdaf5742d7e699022c64ecd/mscorlib/system/io/path.cs#L999
func getFileName(path string) string {
	if path != "" {
		length := len(path)
		for i := length - 1; i >= 0; i-- {
			ch := path[i]
			// backslash ("\"), slash ("/"), or colon (":")
			if ch == '\\' || ch == '/' || ch == ':' {
				return path[i+1:]
			}
		}
	}
	return path
}

func appNameFromUrl(url string) string {
	var re = regexp.MustCompile(".json$")
	return re.ReplaceAllString(getFileName(url), "")
}

// paths

// fname retuns filename
func fname(path string) string {
	return getFileName(path)
}

// stripExt returns fname without extension
func stripExt(fname string) string {
	return strings.TrimSuffix(fname, filepath.Ext(fname))
}

// stripFilename returns path without the filename
func stripFilename(path string) string {
	return strings.TrimSuffix(path, fname(path))
}

// stripFragment removes the fragment part of a uri
func stripFragment(uri string) string {
	u, _ := url.Parse(uri)
	return strings.TrimSuffix(uri, "#"+u.Fragment)
}

// appList convert list of apps to a map of app->isGlobal map
func appList(apps []string, isGlobal bool) map[string]bool {
	m := map[string]bool{}
	if len(apps) == 0 {
		return m
	}
	for _, app := range apps {
		m[app] = isGlobal
	}
	return m
}

func parseApp(app string) (string, string, string) {
	re := regexp.MustCompile("(?:(?P<bucket>[a-zA-Z0-9-]+)/)?(?P<app>.*.json$|[a-zA-Z0-9-_.]+)(?:@(?P<version>.*))?")

	if re.MatchString(app) {
		match := re.FindStringSubmatch(app)
		result := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}

		return result["app"], result["bucket"], result["version"]
	}

	return app, "", ""
}

func showApp(app, bucket, version string) string {
	if bucket != "" {
		app = bucket + "/" + app
	}
	if version != "" {
		app = app + "@" + version
	}
	return app
}

// core envs

func getScoopDir() string {
	if os.Getenv("SCOOP") != "" {
		return os.Getenv("SCOOP")
	} else if getConfig("rootPath") != "" {
		return getConfig("rootPath")
	} else {
		return filepath.Join(os.Getenv("USERPROFILE"), "SCOOP")
	}
}

func getGlobalDir() string {
	if os.Getenv("SCOOP_GLOBAL") != "" {
		return os.Getenv("SCOOP_GLOBAL")
	} else if getConfig("globalPath") != "" {
		return getConfig("globalPath")
	} else {
		return filepath.Join(os.Getenv("ProgramData"), "SCOOP")
	}
}

func getCacheDir() string {
	if os.Getenv("SCOOP_CACHE") != "" {
		return os.Getenv("SCOOP_CACHE")
	} else if getConfig("cachePath") != "" {
		return getConfig("cachePath")
	} else {
		return filepath.Join(ScoopDir, "cache")
	}
}

func getConfigHome() string {
	if os.Getenv("XDG_CONFIG_HOME") != "" {
		return os.Getenv("XDG_CONFIG_HOME")
	} else {
		return filepath.Join(os.Getenv("USERPROFILE"), ".config")
	}
}

// ScoopDir represents for Scoop base dir
var ScoopDir = getScoopDir()

// GlobalDir represents for Scoop Global dir
var GlobalDir = getGlobalDir()

// CacheDir represents for downloaded cache dir
var CacheDir = getCacheDir()

// ConfigFile returns the path of scoop config file
var ConfigFile = filepath.Join(getConfigHome(), "scoop", "config.json")

// scoopConfig stands for scoop config file content
var scoopConfig, _ = loadConfig(ConfigFile)
