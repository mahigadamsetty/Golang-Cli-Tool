package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		fmt.Println(err)
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func WriteToFile(File string, content string) error {
	exists, err := exists(File)
	if err != nil {
		return err
	}
	var file *os.File
	if !exists {
		fmt.Println(exists)
		file, err = os.Create(File)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(exists)
		file, err = os.OpenFile(File, os.O_RDWR|os.O_APPEND, 0660)
		if err != nil {
			log.Fatal(err)
		}
	}

	len, err := file.WriteString(content + "\n")
	if err != nil {
		return err
	}
	fmt.Printf("length : %d", len)

	return nil
}

func ReadFromFile(File string) error {
	data, err := ioutil.ReadFile(File)
	if err != nil {
		return err
	}

	fmt.Printf("\nList of movies:\n %s", data)

	return nil
}
