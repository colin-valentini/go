package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

type InitCmd struct {
	Id   int    `arg:"" name:"id" required:"" help:"The LeetCode question identifier." type:""`
	Slug string `arg:"" name:"slug" required:"" help:"The LeetCode question title slug." type:""`
}

type TemplateData struct {
	Id                         int
	Slug, SlugCamel, SlugTitle string
}

func (ic *InitCmd) Run() error {
	log.Printf("Running LeetCode init command for id=%d, slug=%s\n", ic.Id, ic.Slug)
	templateFiles, err := templateFiles()
	if err != nil {
		return fmt.Errorf("reading template files: %s\n", err)
	}

	targetDir := filepath.Join(leetCodeDir(), fmt.Sprintf("%d_%s", ic.Id, ic.Slug))
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("creating leetcode target dir: %s\n", err)
	}
	log.Printf("Created target directory at %s\n", targetDir)

	templateData := ic.templateData()
	for _, tempFile := range templateFiles {
		fName := fmt.Sprintf("%s.go", filepath.Base(tempFile))
		f, err := os.Create(filepath.Join(targetDir, fName))
		if err != nil {
			return fmt.Errorf("creating file %s: %s\n", fName, err)
		}
		defer f.Close()
		temp, err := template.ParseFiles(tempFile)
		if err != nil {
			return fmt.Errorf("parsing template file for %s: %s\n", fName, err)
		}
		if temp.Execute(f, templateData); err != nil {
			return fmt.Errorf("executing template file for %s: %s\n", fName, err)
		}
		log.Printf("Generated file %s\n", fName)
	}
	log.Println("Done")
	return nil
}

func (ic *InitCmd) templateData() TemplateData {
	return TemplateData{
		Id:        ic.Id,
		Slug:      ic.Slug,
		SlugCamel: transformCase(ic.Slug, camelCase),
		SlugTitle: transformCase(ic.Slug, titleCase),
	}
}

func templateFiles() ([]string, error) {
	tDir := templateDir()
	entries, err := os.ReadDir(tDir)
	if err != nil {
		return nil, fmt.Errorf("reading directory %s: \n", err)
	}
	files := make([]string, 0, len(entries))
	for _, entry := range entries {
		// Skip any subdirectories
		if entry.IsDir() {
			continue
		}
		files = append(files, filepath.Join(tDir, entry.Name()))
	}
	return files, nil
}

func templateDir() string {
	return filepath.Join(basepath, "templates")
}

func leetCodeDir() string {
	return filepath.Join(basepath, "../../practice/leetcode")
}

type caseStyle int

const (
	titleCase caseStyle = iota
	camelCase
)

func transformCase(slug string, style caseStyle) string {
	elems := strings.Split(slug, "-")
	for i := range elems {
		if style == camelCase && i == 0 {
			elems[i] = strings.ToLower(elems[i])
			continue
		}
		elem := []rune(elems[i])
		elems[i] = strings.ToUpper(string(elem[0])) + string(elem[1:])
	}
	return strings.Join(elems, "")
}
