package base

import (
	"os"
)

type CloudStorageBase struct {
}

func (c CloudStorageBase) getRootPath() string {
	return os.Getenv("HOME_URL")
}

func (c CloudStorageBase) GetKeyFile() string {
	pwd := os.Getenv("HOME_URL")
	return pwd + "/config/cloud_storage/" + "gcs.json"
}

func (c CloudStorageBase) GetStorageBucket() string {
	return os.Getenv("GOOGLE_CLOUD_STORAGE_BUCKET")
}

func (c CloudStorageBase) GetProjectId() string {
	return os.Getenv("GOOGLE_CLOUD_PROJECT_ID")
}

func (c CloudStorageBase) GetStoragePathPrefix() string {
	return os.Getenv("GOOGLE_CLOUD_STORAGE_PATH_PREFIX")
}
