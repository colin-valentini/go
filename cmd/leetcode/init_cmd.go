package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type InitCmd struct {
	Id   int    `arg:"" name:"id" required:"" help:"The LeetCode question identifier." type:""`
	Slug string `arg:"" name:"slug" required:"" help:"The LeetCode question title slug." type:""`
	Dir  string `arg:"" name:"dir" required:"" help:"The directory where the files should be created." type:"path"`
}

type TemplateData struct {
	Id                         int
	Slug, SlugCamel, SlugTitle string
}

func (ic *InitCmd) Run() error {
	fmt.Printf("id: %d, slug: %q, dir: %s\n", ic.Id, ic.Slug, ic.Dir)
	folder := fmt.Sprintf("%s/%d_%s", ic.Dir, ic.Id, ic.Slug)
	if err := os.MkdirAll(folder, 0755); err != nil {
		return fmt.Errorf("creating leetcode folder: %s\n", err)
	}
	tmpDir, err := filepath.Abs("./cmd/leetcode/templates/")
	if err != nil {
		return fmt.Errorf("getting path to templates directory: %s\n", err)
	}
	tmpData := TemplateData{
		Id:        ic.Id,
		Slug:      ic.Slug,
		SlugCamel: camelCase(ic.Slug),
		SlugTitle: titleCase(ic.Slug),
	}
	if err := genSolutionGo(folder, tmpData, tmpDir); err != nil {
		return err
	}
	if err := genSolutionTestGo(folder, tmpData, tmpDir); err != nil {
		return err
	}
	return nil
}

func genSolutionGo(folder string, tmpData TemplateData, tmpDir string) error {
	solutionFile, err := os.Create(fmt.Sprintf("%s/solution.go", folder))
	if err != nil {
		return fmt.Errorf("creating solution.go file: %s\n", err)
	}
	defer solutionFile.Close()
	solutionTmp, err := template.ParseFiles(tmpDir + "/solution")
	if err != nil {
		return fmt.Errorf("parsing solution template file: %s\n", err)
	}
	if err := solutionTmp.Execute(solutionFile, tmpData); err != nil {
		return fmt.Errorf("executing solution template: %s\n", err)
	}
	return nil
}

func genSolutionTestGo(folder string, tmpData TemplateData, tmpDir string) error {
	solutionTestFile, err := os.Create(fmt.Sprintf("%s/solution_test.go", folder))
	if err != nil {
		return fmt.Errorf("creating solution_test.go file: %s\n", err)
	}
	defer solutionTestFile.Close()
	solutionTmp, err := template.ParseFiles(tmpDir + "/solution_test")
	if err != nil {
		return fmt.Errorf("parsing solution_test template file: %s\n", err)
	}
	if err := solutionTmp.Execute(solutionTestFile, tmpData); err != nil {
		return fmt.Errorf("executing solution_test template: %s\n", err)
	}
	return nil
}

func titleCase(slug string) string {
	elems := strings.Split(slug, "-")
	for i := range elems {
		elem := []rune(elems[i])
		elems[i] = strings.ToUpper(string(elem[0])) + string(elem[1:])
	}
	return strings.Join(elems, "")
}

func camelCase(slug string) string {
	elems := strings.Split(slug, "-")
	for i := range elems {
		if i == 0 {
			elems[i] = strings.ToLower(elems[i])
			continue
		}
		elem := []rune(elems[i])
		elems[i] = strings.ToUpper(string(elem[0])) + string(elem[1:])
	}
	return strings.Join(elems, "")
}
