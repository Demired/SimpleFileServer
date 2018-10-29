package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
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

	var wanIPURL = "https://myip.ipip.net/s"

	wanIP := httpGet(wanIPURL)

	IntranetIP := getIntranetIP()

	fmt.Printf("listen port is %d\n", port)
	fmt.Println("----------------------------------------------")
	fmt.Printf("| loc address : http://127.0.0.1:%d\n", port)
	fmt.Printf("| lan address : http://%s:%d\n", IntranetIP, port)
	fmt.Printf("| wan address : http://%s:%d\n", strings.TrimSpace(wanIP), port)
	fmt.Println("----------------------------------------------")

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(dir)))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getIntranetIP() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

func httpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return string(body)
}
