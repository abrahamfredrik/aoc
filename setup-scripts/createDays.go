package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	for i := 5; i <= 25; i++ {
		newpath := filepath.Join("../2023", "day")
		if i < 10 {
			newpath = newpath + "0"
		}
		newpath = newpath + fmt.Sprint(i)
		os.MkdirAll(newpath, os.ModePerm)

		copyFile("template/main.go", newpath+"/main.go")
		copyFile("template/example.txt", newpath+"/example.txt")
		copyFile("template/dayInput.txt", newpath+"/dayInput.txt")
	}

}

func copyFile(src string, dst string) {
	data, _ := ioutil.ReadFile(src)
	ioutil.WriteFile(dst, data, 0644)
}
