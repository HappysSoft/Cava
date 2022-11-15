package model

type Build struct {
	version         string   `yaml:"cava"`
	projectName     string   `yaml:"project"`
	mainClass       string   `yaml:"main"`
	includedPath    []string `yaml:"include path"`
	includedLibrary []string `yaml:"include library"`
	outputPath      string   `yaml:"output path"`
	outputName      string   `yaml:"output"`
	outputType      int      `yaml:"type"`
}
