package chopsticks

import (
	"io/ioutil"
	"sort"
	"strings"
)

// LatestVersion get the version from bucket
func LatestVersion(app, bucket string) (string, error) {
	mf, err := GetManifest(app, bucket)
	if err != nil {
		return "", err
	}

	return mf.Version, nil
}

// LatestVersionFromUrl get the latest version from url
func LatestVersionFromUrl(url string) (string, error) {
	mf := GetUrlManifest(url)

	return mf.Version, nil
}

// CurrentVersion
func CurrentVersion(app string, isGlobal bool) (string, error) {
	v, err := getVersions(app, isGlobal)
	if err != nil {
		return "", err
	}

	v = sortVersion(v)

	return v[len(v)-1], nil
}

// getVersions return a string array of app's all versions
// exclude the `current` version
func getVersions(app string, isGlobal bool) ([]string, error) {
	appdir := appDir(app, isGlobal)

	dirs, err := ioutil.ReadDir(appdir)
	if err != nil {
		return []string{}, err
	}

	var vs []string
	for _, dir := range dirs {
		if dir.IsDir() {
			vs = append(vs, dir.Name())
		}
	}

	return vs, nil
}

type versions []string

func (v versions) Len() int {
	return len(v)
}
func (v versions) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v versions) Less(i, j int) bool {
	toVersionArray := func(v string) []string {
		return strings.FieldsFunc(v, func(r rune) bool {
			return r == '.' || r == '-'
		})
	}

	a := toVersionArray(v[i])
	b := toVersionArray(v[j])

	for i, v := range a {
		if i > len(b) {
			return false
		}

		if v > b[i] {
			return false
		}
		if v < b[i] {
			return true
		}
	}

	if len(b) > len(a) {
		return true
	}

	return false
}

func sortVersion(vs []string) []string {
	sort.Sort(versions(vs))

	return vs
}

// compareVersions compare two version string a and b
// if a is newer than b, return 1
// if a is equal to b, return 0
// if a is older than b, return -1
func compareVersions(a, b string) int {
	toVersionArray := func(v string) []string {
		return strings.FieldsFunc(v, func(r rune) bool {
			return r == '.' || r == '-'
		})
	}

	verA := toVersionArray(a)
	verB := toVersionArray(b)

	for i, v := range verA {
		if i > len(b) {
			return 1
		}

		if v > verB[i] {
			return 1
		}
		if v < verB[i] {
			return -1
		}
	}

	if len(b) > len(a) {
		return -1
	}

	return 0
}
