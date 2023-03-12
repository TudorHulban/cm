package main

import (
	"io/fs"
	"strings"
	"test/app/apperrors"
	"test/helpers"
)

func loadFSEntries() ([]fs.DirEntry, error) {
	fsEntries, errRead := fs.ReadDir(_fs, _folderBackupFiles)
	if errRead != nil {
		return nil, &apperrors.ErrInfrastructure{
			Caller:  "loadFSEntries",
			Calling: "fs.ReadDir",
			Issue:   errRead,
		}
	}

	return fsEntries, nil
}

// TODO; move to one function with prefix param.
func loadVarsFiles(fsEntries []fs.DirEntry) (map[string][]byte, error) {
	res := make(map[string][]byte)

	for _, entry := range fsEntries {
		if entry.IsDir() {
			continue
		}

		if !strings.HasPrefix(entry.Name(), _prefixFilesVars) {
			continue
		}

		content, errRead := _fs.ReadFile(_folderBackupFiles + "/" + entry.Name())
		if errRead != nil {
			return nil, errRead
		}

		res[helpers.FileNameWOExtension(entry.Name())] = content
	}

	return res, nil
}

func loadServiceFiles(fsEntries []fs.DirEntry) (map[string][]byte, error) {
	res := make(map[string][]byte)

	for _, entry := range fsEntries {
		if entry.IsDir() {
			continue
		}

		if !strings.HasPrefix(entry.Name(), _prefixFilesServices) {
			continue
		}

		content, errRead := _fs.ReadFile(_folderBackupFiles + "/" + entry.Name())
		if errRead != nil {
			return nil, errRead
		}

		res[helpers.FileNameWOExtension(entry.Name())] = content
	}

	return res, nil
}

func loadTargetFiles(fsEntries []fs.DirEntry) (map[string][]byte, error) {
	res := make(map[string][]byte)

	for _, entry := range fsEntries {
		if entry.IsDir() {
			continue
		}

		if !strings.HasPrefix(entry.Name(), _prefixFilesTargets) {
			continue
		}

		content, errRead := _fs.ReadFile(_folderBackupFiles + "/" + entry.Name())
		if errRead != nil {
			return nil, errRead
		}

		res[helpers.FileNameWOExtension(entry.Name())] = content
	}

	return res, nil
}
