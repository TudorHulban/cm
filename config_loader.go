package main

import (
	"io/fs"
	"test/helpers"
)

func loadConfigFiles() (map[string][]byte, error) {
	fsEntries, errRead := fs.ReadDir(_fs, _folderBackupFile)
	if errRead != nil {
		return nil, errRead
	}

	res := make(map[string][]byte)

	for _, entry := range fsEntries {
		if entry.IsDir() {
			continue
		}

		content, errRead := _fs.ReadFile(_folderBackupFile + "/" + entry.Name())
		if errRead != nil {
			return nil, errRead
		}

		res[helpers.FileNameWOExtension(entry.Name())] = content
	}

	return res, nil
}
