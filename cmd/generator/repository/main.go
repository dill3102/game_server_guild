package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	ctrl := getParam()
	ctrl.Print()
	if err := os.MkdirAll(ctrl.OutputPath, 0755); err != nil {
		panic(err)
	}

	original := GetTemplateFile(ctrl.TemplatePath)
	models := GetFiles(ctrl.ModelPath)
	for _, model := range models {
		modelName, implName := GenName(strings.Split(model, "_"))
		replaced := strings.Replace(original, "{ModelName}", modelName, -1)
		replaced = strings.Replace(replaced, "{ImplName}", implName, -1)
		writeFile(ctrl.OutputPath+"/"+model+"_repository.go", replaced)
	}
}

func GenName(names []string) (string, string) {

	firstUpperName := ""
	firstLowerName := ""

	for i, n := range names {
		if len(n) == 0 {
			continue
		}

		runes := []rune(n)
		runes[0] = []rune(strings.ToUpper(string(runes[0])))[0]
		firstUpperName += string(runes)
		firstLowerName += string(runes)
		if i == 0 {
			firstLowerName = strings.ToLower(string(runes))
		}
	}

	return firstUpperName, firstLowerName
}

func writeFile(filePath, file string) {
	if err := os.WriteFile(filePath, []byte(file), 0664); err != nil {
		panic(err)
	}
}

func GetTemplateFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("ファイル読み込みエラー:", err)
		return ""
	}

	return string(data)
}

func GetFiles(path string) []string {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("ディレクトリ読み込みエラー:", err)
		return []string{}
	}

	ret := []string{}
	for _, e := range entries {
		ret = append(ret, strings.Split(e.Name(), ".")[0])
	}
	return ret
}

type DirectoryCtrl struct {
	ModelPath    string
	TemplatePath string
	OutputPath   string
}

func getParam() *DirectoryCtrl {
	if len(os.Args) < 4 {
		panic("使い方: go run main.go <model path>, <template path>, <Output Path>")
		return nil
	}
	return &DirectoryCtrl{
		ModelPath:    os.Args[1],
		TemplatePath: os.Args[2],
		OutputPath:   os.Args[3],
	}
}

func (m *DirectoryCtrl) Print() {
	fmt.Println("ModelPath: ", m.ModelPath)
	fmt.Println("TemplatePath: ", m.TemplatePath)
}
