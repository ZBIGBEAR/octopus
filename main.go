package main

import (
	"flag"
	"fmt"
)

var (
	path string
)

func init() {
	flag.StringVar(&path, "path", ".", "请输入入口文件或者文件夹:-path")
}

func main(){
	flag.Parse()
	fmt.Println("filePath:",path)
}