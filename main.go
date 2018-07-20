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

var build string

var version string

func main() {
	for _, p := range os.Args {
		switch p {
		case "-h":
			fmt.Println("Usage of SimpleFileServer [-v] [-h] [--port=8000] [--dir=.]\n  -v\n        version number\n  --port int\n        listen http port (default 8000)\n  --dir string\n        listen file path (default .)")
			return
		case "-v":
			fmt.Printf("SimpleFileServer for Linux/Mac/BSD, Version %s, Build %s.\n", version, build)
			return
		}
		if strings.Contains(p, "--port=") {
			args := strings.Split(p, "=")
			if len(args) == 2 {
				i, _ := strconv.Atoi(args[1])
				port = i
			}
		} else {
			port = 8000
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
		} else {
			dir = "."
		}
	}
	fmt.Printf("listen port is %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(dir)))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
