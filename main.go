package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var projectName string
	//引数処理
	switch len(os.Args) {
	case 2:
		projectName = os.Args[1]
	case 1:
		log.Fatalln("プロジェクト名を指定してください")
	default:
		log.Fatalln("プロジェクト名は複数指定することはできません")
	}
	project := &Project{name: projectName}
	if err := project.create(); err != nil {
		panic(err)
	}
	fmt.Printf("プロジェクト %v を作成しました\n", projectName)
	fmt.Printf("./%v\n", projectName)
	fmt.Println("  ├── README.md")
	fmt.Println("  └── main.go")
}
