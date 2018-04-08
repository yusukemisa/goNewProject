package main

import (
	"fmt"
	"os"
	"text/template"
)

// Project 構造体
type Project struct {
	name string
}

func (p *Project) create() error {
	//プロジェクトディレクトリ作成
	if err := os.Mkdir(p.name, 0755); err != nil {
		fmt.Printf("プロジェクトディレクトリの作成に失敗しました。\n%v", err)
		return err
	}

	//README作成
	if err := createFileFromTemplate(p.name, "README.md"); err != nil {
		return err
	}
	//main.go作成
	return createFileFromTemplate(p.name, "main.go")
}

func createFileFromTemplate(projectName string, fileName string) error {
	file, err := os.Create(projectName + "/" + fileName)
	defer file.Close()
	if err != nil {
		fmt.Printf("%vの作成に失敗しました\n%v", fileName, err)
		return err
	}
	t := template.Must(template.ParseFiles("templates/" + fileName))
	t.Execute(file, projectName)
	return nil
}
