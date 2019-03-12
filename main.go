package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func sendExcel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=README.md")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	data, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Fatal("读取文件错误")
	}
	http.ServeContent(w, r, "README.md", time.Now(), bytes.NewReader(data))
}

func testR(w http.ResponseWriter, r *http.Request) {
	fmt.Println("xxxx")
}

func main() {
	http.HandleFunc("/test", testR)
	http.HandleFunc("/excel", sendExcel)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("Some Error Happends")
	}
}