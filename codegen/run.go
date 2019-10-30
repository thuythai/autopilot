package codegen

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func Run(dir string) error {
	config := filepath.Join(dir, "autopilot.yaml")

	project, err := Load(config)
	if err != nil {
		return err
	}

	files, err := Generate(project)
	if err != nil {
		return err
	}

	for name, content := range files {
		if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
			return err
		}
		if err := ioutil.WriteFile(name, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}
