package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"mxtransporter/pkg/errors"
	"os"
)

func init() {
	m := godotenv.Load()
	if m != nil {
		fmt.Println("[Warning] If this environment is local machine, you have to create .env file, and set env variables with reference to .env.template .")
	}
}

func FetchPersistentVolumeDir() (string, error) {
	pvDir, pvDirExistence := os.LookupEnv("PERSISTENT_VOLUME_DIR")
	if pvDirExistence == false {
		return "", errors.InternalServerErrorEnvGet.New("PERSISTENT_VOLUME_DIR is not existed in environment variables")
	}
	return pvDir, nil
}

func FetchExportDestination() (string, error) {
	exportDestination, exportDestinationExistence := os.LookupEnv("EXPORT_DESTINATION")
	if exportDestinationExistence == false {
		return "", errors.InternalServerErrorEnvGet.New("EXPORT_DESTINATION is not existed in environment variables")
	}
	return exportDestination, nil
}

func FetchGcpProject() (string, error) {
	projectID, projectIDExistence := os.LookupEnv("PROJECT_NAME_TO_EXPORT_CHANGE_STREAMS")
	if projectIDExistence == false {
		return "", errors.InternalServerErrorEnvGet.New("PROJECT_NAME_TO_EXPORT_CHANGE_STREAMS is not existed in environment variables")
	}
	return projectID, nil
}
