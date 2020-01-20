package Create_file

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func readtxt() []os.FileInfo {
	files, err := ioutil.ReadDir("./TXT/")
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func openreadtext() {
	files := readtxt()
	for _, file := range files {
		content, err := ioutil.ReadFile(fmt.Sprint("TXT/", file))
		if err != nil {
			log.Fatal(err)
		}
		re := regexp.MustCompile(`[Atiwan]\w+/g`)
		cresultName, _ := regexp.MatchString(content)
		retur
	}

	//fmt.Printf("File contents: %s", content)
}
