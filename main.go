package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var dir string

var port int

func main() {
	// fmt.Println(os.Args)
	// fmt.Println(strings.Contains("123a", "123"))
	for _, p := range os.Args {
		switch p {
		case "-h":
			fmt.Println("print help word")
			return
		case "-v":
			fmt.Println("print version word")
			return
		}
		if strings.Contains(p, "--port=") {
			args := strings.Split(p, "=")
			if len(args) == 2 {
				i, _ := strconv.Atoi(args[1])
				port = i
			}

		}
		if strings.Contains(p, "--dir=") {
			args := strings.Split(p, "=")
			if len(args) == 2 {
				stat, err := os.Stat(args[1]) //os.Stat获取文件信息
				if err == nil && stat.IsDir() {
					dir = args[1]
					continue
				}
				fmt.Println("dir does not exist")
				return
			}
		}
	}
	if port == 0 {
		port = 8000
	}
	fmt.Printf("listen port is %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(dir)))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
