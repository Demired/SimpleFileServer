package main

import (
	"fmt"
	"log"
	"testing"
)

func TestHttpGet(t *testing.T) {
	var url = "https://myip.ipip.net/s"
	res := httpGet(url)
	fmt.Println(len(res) > 1)
	if len(res) < 1 {
		log.Fatal("HttpGet失败")
	}
}
