package chopsticks

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

var BucketsDir = filepath.Join(ScoopDir, "buckets")

func FindBucketDirectory(name string) string {
	bucket := filepath.Join(BucketsDir, name)

	if f, err := os.Stat(filepath.Join(bucket, "bucket")); !os.IsNotExist(err) && f.IsDir() {
		bucket = filepath.Join(bucket, "bucket")
	}

	return bucket
}

func GetLocalBuckets() ([]string, error) {
	dirs, err := ioutil.ReadDir(BucketsDir)
	if err != nil {
		return []string{}, err
	}

	var buckets []string
	for _, dir := range dirs {
		if dir.IsDir() {
			buckets = append(buckets, dir.Name())
		}
	}

	return buckets, nil
}

func findManifestInBucket(app, bucket string) {
	var mf Manifest
	if bucket != "" {
		mf = manifes
	}
}

func AddBucket(name, repo string) {

}

func RmBucket(name string) {

}

func FindBucketDir(name string) (string, error) {
	bucket := filepath.Join(BucketsDir, name)

	stat, err := os.Stat(bucket)

	if err == nil && stat.IsDir() {
		return filepath.Join(bucket, "bucket"), nil
	} else {
		return bucket, err
	}
}
