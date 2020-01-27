package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type file struct {
	Sizes int64
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
func FindFileFromExtension(extension []string, dir string, files *[]file) {
	fs, err := ioutil.ReadDir(dir)
	if err == nil {
		for _, f := range fs {
			for _, ex := range extension {
				if strings.HasSuffix(f.Name(), ex) {
					if (time.Now().Unix() - f.ModTime().Unix()) < 2592000 {
						*files = append(*files, file{
							Sizes: f.Size(),
							path:  dir,
							name:  f.Name(),
						})

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
	start := time.Now()
	drives := getDrives()
	files := []file{}
	for _, d := range drives {
		FindFileFromExtension([]string{".jpg", ".gif", ".png"}, d, &files)
	}
	fmt.Println(len(files))
	fmt.Println(time.Since(start))
	/*	for _, i := range files {
		if i.Sizes < 1000000 {
			fmt.Println(len(files))
			fmt.Println(i.Sizes, i.name, i.path)
		}
		if i.Sizes > 1000000 {
			fmt.Println(len(files))
			fmt.Println(i.Sizes, i.name, i.path)
		}*/
	//}

}
