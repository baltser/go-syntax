package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
	}
	file := arguments[1]

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	//fmt.Println(pathSplit)
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)

		//Сущевствует ли файл
		fileInfo, err := os.Stat(fullPath)

		if err == nil {
			mode := fileInfo.Mode()

			//Это обычный файл?
			if mode.IsRegular() {
				// Является ли он исполняемым?
				if mode&0111 != 0 {
					fmt.Println(fullPath)
					return
				}
			}
		}

	}
}
