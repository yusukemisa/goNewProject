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
	fmt.Printf("プロジェクトディレクトリ %v を作成しました\n", p.name)
	//README作成
	if err := createReadMe(p.name); err != nil {
		return err
	}
	//main.go作成
	return createMain(p.name)
}

func createReadMe(name string) error {
	file, err := os.Create(name + "/README.md")
	defer file.Close()
	if err != nil {
		fmt.Printf("README.mdの作成に失敗しました\n%v", err)
		return err
	}
	t := template.Must(template.ParseFiles("templates/README"))
	t.Execute(file, name)
	return nil
}

func createMain(name string) error {
	file, err := os.Create(name + "/main.go")
	defer file.Close()
	if err != nil {
		fmt.Printf("main.goの作成に失敗しました\n%v", err)
		return err
	}
	t := template.Must(template.ParseFiles("templates/Main"))
	t.Execute(file, name)
	return nil
}
