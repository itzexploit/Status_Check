package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func banner() {
	banner := `
   (\_/) %vCreated By Exploit%v
   (0.0) %vVersion %v1.0%v
   (\>❤  %vdiscord.gg/pytholearn %v| %vdiscord.gg/arabic%v
`
	fmt.Printf(banner, Green, Reset, Red, Blue, Reset, Cyan, Red, Gray, Reset)
}
func main() {
	banner()
	if len(os.Args) < 2 {
		fmt.Printf("Please Give FileName Example main.go example.txt")
		return
	}

	fileContent, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	urls := strings.Fields(string(fileContent))

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error :", err)
			continue
		}
		defer resp.Body.Close()
		fmt.Println(White+"\n - URL :", Magenta+url, White+": Status Code :", Green, resp.StatusCode, Reset, "\n")
	}
}
