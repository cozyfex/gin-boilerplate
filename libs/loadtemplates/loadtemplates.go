package loadtemplates

import (
	"fmt"
	"io/ioutil"
)

type File struct {
	name string
	file string
}

func GetFiles(dir string) []string {
	var result []string

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return result
	}

	for _, file := range files {
		filename := dir + "/" + file.Name()
		if file.IsDir() {
			f := GetFiles(filename)
			result = append(result, f...)
		} else {
			fmt.Println(filename)
			result = append(result, filename)
		}
	}

	return result
}
