package helpers

import (
	"path"
	"path/filepath"

	"promotions-app/config"

	guuid "github.com/google/uuid"
)

func UniqueTmpPath(fileName string) string {
	config := config.GetConfig()

	uName := guuid.New().String() + filepath.Ext(fileName)

	return path.Join(config.GetString("tmpFolder"), uName)
}
