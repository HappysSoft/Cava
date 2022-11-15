package util

import (
	"Cava/model"
	yaml "gopkg.in/yaml.v3"

	"os"
)

func ReadSourceCodeFile(path string) error {
	return nil
}

func ReadBuildFile(path string) (*model.Build, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var buildFile model.Build
	err = yaml.Unmarshal(file, &buildFile)
	if err != nil {
		return nil, err
	}
	return &buildFile, err
}
