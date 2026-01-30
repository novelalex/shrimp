package fileutil

import (
	"maps"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func ExecutablesInPath(path string) (map[string]string, error) {
	pathString := path
	pathDirs := strings.Split(pathString, string(os.PathListSeparator))

	executablesInPath := make(map[string]string)
	for _, dir := range pathDirs {
		executablesInDir, err := ExecutablesInDir(dir)
		if err != nil {
			return nil, err
		}
		maps.Copy(executablesInPath, executablesInDir)
	}

	return executablesInPath, nil
}

func ExecutablesInDir(dirPath string) (map[string]string, error) {
	executables := make(map[string]string)

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return executables, nil
		//return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			continue
		}

		name := entry.Name()

		if runtime.GOOS == "windows" {
			if isWindowsExecutable(entry.Name()) {
				executables[name] = filepath.Join(dirPath, name)
			}
		} else {
			// 0111 masks x bits
			if info.Mode()&0111 != 0 {
				executables[name] = filepath.Join(dirPath, name)
			}
		}
	}

	return executables, nil
}

func isWindowsExecutable(name string) bool {
	ext := filepath.Ext(name)
	switch ext {
	case ".exe", ".bat", ".cmd", "ps1":
		return true
	}
	return false
}
