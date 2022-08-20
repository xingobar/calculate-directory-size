package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Please enter the absolute path:")

	reader := bufio.NewReader(os.Stdin)

	directory, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Get wrong! Please try again.", err)
	}

	directory = strings.TrimSuffix(directory, "\n")

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		fmt.Printf("%s directory is not exists. \n", directory)
		return
	}

	if directoryInfo, err := os.Stat(directory); err != nil || !directoryInfo.IsDir() {
		fmt.Println("Get directory info error or it is not directory .")
		return
	}

	var size int64

	err = filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			size += info.Size()
		}

		return err
	})

	if err != nil {
		fmt.Println("file path walk recursive wrong! ", err)

		return
	}

	mb := float64(size) / 1024 / 1204

	fmt.Printf("Directory Size is %f MB \n", mb)
}