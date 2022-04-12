package utils

import (
	"mngr/models"
	"os"
	"path"
	"strings"
)

func createDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateRequiredDirectories(config *models.Config) {
	// Create HLS stream folder
	stream := GetStreamPath(config)
	createDirIfNotExist(stream)

	// Create record folder
	record := GetRecordPath(config)
	createDirIfNotExist(record)

	// Create object detection folder
	od := GetOdPath(config)
	createDirIfNotExist(od)

	// Create facial recognition folder
	fr := GetFrPath(config)
	createDirIfNotExist(fr)
}

func CreateSourceDefaultDirectories(config *models.Config, sourceId string) {
	// Create HLS stream folder for the source
	stream := GetStreamPath(config)
	createDirIfNotExist(path.Join(stream, sourceId))

	// Create record folder for the source
	record := GetRecordPath(config)
	createDirIfNotExist(path.Join(record, sourceId))
	//and also short video clips folder
	createDirIfNotExist(path.Join(record, sourceId, "temp"))

	// Create object detection folder for the source
	od := GetOdPath(config)
	createDirIfNotExist(path.Join(od, sourceId))
	createDirIfNotExist(path.Join(od, sourceId, "data"))
	createDirIfNotExist(path.Join(od, sourceId, "images"))
	createDirIfNotExist(path.Join(od, sourceId, "videos"))

	// Create facial recognition folder for the source
	fr := GetFrPath(config)
	createDirIfNotExist(path.Join(fr, sourceId))
	createDirIfNotExist(path.Join(fr, sourceId, "data"))
	createDirIfNotExist(path.Join(fr, sourceId, "images"))
	createDirIfNotExist(path.Join(fr, sourceId, "videos"))
}

func GetStreamPath(config *models.Config) string {
	return path.Join(config.General.RootFolderPath, "stream")
}

func GetRecordPath(config *models.Config) string {
	return path.Join(config.General.RootFolderPath, "record")
}

func GetOdPath(config *models.Config) string {
	return path.Join(config.General.RootFolderPath, "od")
}

func GetFrPath(config *models.Config) string {
	return path.Join(config.General.RootFolderPath, "fr")
}

func SetHlsPath(config *models.Config, s *models.StreamModel) {
	s.HlsOutputPath = strings.Replace(s.HlsOutputPath, path.Join(config.General.RootFolderPath, "stream"), "", -1)
}

func GetOdImagesPathBySourceId(config *models.Config, sourceId string) string {
	return path.Join(GetOdPath(config), sourceId, "images")
}

func GetOdDataPathBySourceId(config *models.Config, sourceId string) string {
	return path.Join(GetOdPath(config), sourceId, "data")
}