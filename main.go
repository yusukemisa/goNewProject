package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var projectName string
	fmt.Println(len(os.Args))
	//引数処理
	switch len(os.Args) {
	case 2:
		projectName = os.Args[1]
	case 1:
		log.Fatalln("プロジェクト名を指定してください")
	default:
		log.Fatalln("プロジェクト名は複数指定することはできません")
	}

	project := Project{name: projectName}
	if err := project.create(); err != nil {
		panic(err)
	}
}
