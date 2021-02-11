package chopsticks

import (
	"regexp"
	"time"
)

func nightlyVersion(date time.Time, quiet bool) string {
	dateStr := date.Format("20060102")
	if !quiet {
		Warn("This is a nightly version. Downloaded files won't be verified.")
	}
	return "nightly-" + dateStr
}

func FindManifest(app, bucket string) {
	var (
		url string
		mf  Manifest
	)

	if regexp.MustCompile(`^(ht|f)tps?://|\\\\`).Match([]byte(app)) {
		url = app
		app = appNameFromUrl(url)
		mf = GetUrlManifest(url)
	} else {
		mf, bucket = FindManifest(app, bucket)
	}
}
