/*
 Copyright 2020 Ekkakak Chimjan(ekkalak.ch@rmuti.ac.th). All Rights Reserved.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.

-----------------------------------------------------------------------------------
Basic Comman Line Programming(CLI) and HTTP with Go, case study : download the image on public instagram
test by run following command
		go run main.go -url="https://www.instagram.com/p/B70TOUvo2QT/" xxx=a b c d
  -- or --
		go run main.go -url="https://www.instagram.com/p/B70TOUvo2QT/"
  -- or --
		go run main.go
*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	argCount := len(os.Args[1:])
	var stringurl string
	fmt.Printf("Total Arguments (excluding program name): %d\n", argCount)

	//Flag package see also :: https://golang.org/pkg/flag/
	flag.StringVar(&stringurl, "url", "https://www.instagram.com/p/B71C8hAJ24C/", "specify path to use.  defaults to https://www.instagram.com/p/B71C8hAJ24C/")
	flag.Parse()
	fmt.Printf("url = %v\n", stringurl)
	fmt.Printf("other args: %+v\n", flag.Args())

	//http package see also :: https://golang.org/pkg/net/http/#Get
	resp, err := http.Get(stringurl)
	if err == nil {
		keyIndex := "meta property=\"og:image\" content=\""
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		stringBody := string(body)
		//fmt.Println(stringBody)
		firstIndex := strings.Index(stringBody, keyIndex)
		if firstIndex > 0 {
			data := stringBody[firstIndex:]
			data = strings.ReplaceAll(data, keyIndex, "")
			lastIndex := strings.Index(data, "\" />")
			imageURL := data[:lastIndex]
			//fmt.Println(imageURL)
			//eg: imageURL = https://instagram.fbkk5-4.fna.fbcdn.net/v/t51.2885-15/e35/p1080x1080/81836190_2522050957905303_5684807475044431256_n.jpg?_nc_ht=instagram.fbkk5-4.fna.fbcdn.net&_nc_cat=103&_nc_ohc=JSncXQB3jhYAX-3SNhX&oh=06e8f8125f9cb28ccce6412d2e52c3e1&oe=5ED8B7D9
			indexExtension := strings.Index(imageURL, "?")
			if indexExtension > 0 {
				filename := imageURL[:indexExtension]
				//now filename = https://instagram.fbkk5-4.fna.fbcdn.net/v/t51.2885-15/e35/p1080x1080/81836190_2522050957905303_5684807475044431256_n.jpg
				index := strings.LastIndex(filename, "/")
				filename = filename[index+1:]
				//now filename = 81836190_2522050957905303_5684807475044431256_n.jpg
				respImg, err := http.Get(imageURL)
				if err == nil {
					defer respImg.Body.Close()
					img, _ := ioutil.ReadAll(respImg.Body)
					err = ioutil.WriteFile(filename, img, 0644)
					if err == nil {
						fmt.Println("save file : \"", filename, "\" successed.")
					}
				}
			}
		}
	}
}
