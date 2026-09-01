package handler

import (
	"fmt"
	"log"
	"net/http"
	"password-manger/templates"
	"strings"
)

// renderTemplateFile renders embedded templates cleanly from binary memory
func renderTemplateFile(w http.ResponseWriter, templatePath string, data interface{}) error {
	name := strings.TrimPrefix(templatePath, "templates/")
	err := templates.Execute(w, name, data)
	if err != nil {
		log.Printf("Error rendering embedded template %s: %v", name, err)
		http.Error(w, fmt.Sprintf("Template Error: %v", err), http.StatusInternalServerError)
		return err
	}
	return nil
}
