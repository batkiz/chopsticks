package chopsticks

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/otiai10/copy"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func GetManifestPath(app, bucket string) (string, error) {
	p := filepath.Join(FindBucketDirectory(bucket), SanitaryPath(app)+".json")

	path, err := filepath.Abs(p)
	if err != nil {
		return "", err
	}
	return path, nil
}

func GetUrlManifest(url string) Manifest {
	r, err := req.Get(url)
	if err != nil {
		Warn(err.Error())
	}

	if r == nil {
		Error(fmt.Sprintf("Can not get content from %v!", url))
		os.Exit(1)
	}

	m, err := UnmarshalManifest(r.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	return m
}

func GetManifest(app, bucket string) (Manifest, error) {
	path, err := GetManifestPath(app, bucket)
	if err != nil {
		return Manifest{}, err
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return Manifest{}, err
	}

	mf, err := UnmarshalManifest(content)
	if err != nil {
		return Manifest{}, err
	}

	return mf, nil
}

func SaveInstalledManifest(app, bucket, dir, url string) {
	if url != "" {
		r, err := req.Get(url)
		if err != nil {
			Warn(err.Error())
		}

		if r == nil {
			Error(fmt.Sprintf("Can not get content from %v!", url))
			os.Exit(1)
		}

		path := filepath.Join(dir, "manifest.json")

		err = r.ToFile(path)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		oriPath, err := GetManifestPath(app, bucket)
		if err != nil {
			log.Fatal(err)
		}

		destPath := filepath.Join(dir, "manifest.json")

		err = copy.Copy(oriPath, destPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetInstalledManifest(app, version string, isGlobal bool) (Manifest, error) {
	v := versionDir(app, version, isGlobal)
	path := filepath.Join(v, "manifest.json")

	mf, err := UnmarshalManifest([]byte(path))
	if err != nil {
		return Manifest{}, err
	}
	return mf, nil
}

func SaveInstallInfo(info InstallInfo, dir string) error {
	path := filepath.Join(dir, "install.json")
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	i, err := UnmarshalInstallInfo([]byte(content))
	if err != nil {
		return err
	}

	newContent, err := i.ToPrettyJson()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetArchitecture() string {
	configArch := getConfig("default-architecture")

	if configArch != "" {
		return configArch
	}

	var arch string
	switch runtime.GOARCH {
	case "386":
		arch = "32bit"
	case "amd64":
		arch = "64bit"
	default:
		panic("Invalid architecture!")
	}

	return arch
}

func GetInstallInfo(app, version string, isGlobal bool) (InstallInfo, error) {
	path := filepath.Join(versionDir(app, version, isGlobal), "install.json")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return InstallInfo{}, err
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return InstallInfo{}, err
	}

	info, err := UnmarshalInstallInfo(content)
	if err != nil {
		return InstallInfo{}, err
	}

	return info, nil
}
