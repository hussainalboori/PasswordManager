package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// renderTemplateFile robustly finds and renders a template file without panicking
func renderTemplateFile(w http.ResponseWriter, relativePath string, data interface{}) error {
	pathsToTry := []string{
		relativePath,
		filepath.Join(".", relativePath),
		filepath.Join("..", relativePath),
	}

	if execPath, err := os.Executable(); err == nil {
		execDir := filepath.Dir(execPath)
		pathsToTry = append(pathsToTry, filepath.Join(execDir, relativePath))
		pathsToTry = append(pathsToTry, filepath.Join(execDir, "..", relativePath))
	}

	var tmpl *template.Template
	var err error
	var foundPath string

	for _, p := range pathsToTry {
		if _, statErr := os.Stat(p); statErr == nil {
			t, parseErr := template.ParseFiles(p)
			if parseErr == nil {
				tmpl = t
				foundPath = p
				break
			}
			err = parseErr
		}
	}

	if tmpl == nil {
		log.Printf("Template render error for %s: %v (searched paths: %v)", relativePath, err, pathsToTry)
		http.Error(w, fmt.Sprintf("Template Error: Could not load %s", relativePath), http.StatusInternalServerError)
		return fmt.Errorf("failed to load template %s: %w", relativePath, err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Template execution error for %s (%s): %v", relativePath, foundPath, err)
		http.Error(w, "Template Execution Error", http.StatusInternalServerError)
		return err
	}

	return nil
}
