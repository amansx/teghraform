package main

import "os"
import "fmt"
import "io/ioutil"

func ReadFeature(name string) string{
	b, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Print(err)
		return ""
	}
	return string(b)
}

func CWD() string{
	if pwd, err := os.Getwd(); err != nil {
		fmt.Println(err)
		return ""
	}else {
		return pwd
	}
}