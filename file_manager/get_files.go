package file_manager

import (
	"os"
	"path"
	"publish/consts"
)

func GetFiles(folder string) ([]string, error) {
	var filesReturn []string
	if folder == consts.BuildOutputsPath {
		getApks, err := GetFiles(path.Join(folder, "apk", "release"))
		if err != nil {
			return nil, err
		}
		filesReturn = append(filesReturn, getApks...)
		getBundles, err := GetFiles(path.Join(folder, "bundle", "play"))
		if err != nil {
			return nil, err
		}
		filesReturn = append(filesReturn, getBundles...)
		return filesReturn, nil
	}
	f, err := os.Open(folder)
	if err != nil {
		return nil, err
	}
	files, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if path.Ext(file.Name()) == ".apk" || path.Ext(file.Name()) == ".aab" {
			filesReturn = append(filesReturn, path.Join(folder, file.Name()))
		}
	}
	return filesReturn, nil
}
