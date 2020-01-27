package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type file struct {
	Sizes int
	path  string
	name  string
}

func getDrives() (r []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			d := string(drive) + ":/"
			r = append(r, d)
			f.Close()
		}
	}
	return
}

//FindFileFromExtension find file
func FindFileFromExtension(extension []string, dir string, files *[]string) {
	fs, err := ioutil.ReadDir(dir)
	if err == nil {
		for _, f := range fs {
			for _, ex := range extension {
				if strings.HasSuffix(f.Name(), ex) {
					if (time.Now().Unix() - f.ModTime().Unix()) < 2592000 {
						if f.Size() > 1000000 {
							*files = append(*files, f.Name())
						}
					}
				}
			}

			if f.IsDir() {
				path := dir + "/" + f.Name()
				FindFileFromExtension(extension, path, files)
			}
		}
	}
}

func main() {
	drives := getDrives()
	files := []string{}
	for _, d := range drives {
		FindFileFromExtension([]string{".jpg", ".gif", ".png"}, d, &files)
	}
	fmt.Println(len(files))
}
