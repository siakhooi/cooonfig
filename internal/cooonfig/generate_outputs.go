package cooonfig

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var openOutputFile = func(path string) (io.WriteCloser, error) {
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
}

func generateOutputs(model map[string]any, outputs []output, templateDir, outputDir string) error {
	for _, output := range outputs {

		templateFile := filepath.Join(templateDir, output.Template)
		outputFile := filepath.Join(outputDir, output.OutFile)

		fmt.Printf("render output: %s => %s\n", templateFile, outputFile)

		t, err := getTemplate(templateFile)
		if err != nil {
			return fmt.Errorf("get template %q: %w", templateFile, err)
		}

		dir := filepath.Dir(outputFile)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		file, err := openOutputFile(outputFile)
		if err != nil {
			return fmt.Errorf("open file %q: %w", outputFile, err)
		}
		err = t.Execute(file, model)
		if err != nil {
			_ = file.Close()
			return fmt.Errorf("execute template %q: %w", templateFile, err)
		}
		if err := file.Close(); err != nil {
			return fmt.Errorf("close file %q: %w", outputFile, err)
		}

	}
	return nil

}
