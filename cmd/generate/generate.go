package main

import (
	"bimg-shellcode-loader/image"
	"os"
)

func main() {
	// 打开shell.bin文件,获取byte数组
	shellcode, _ := os.ReadFile("shell.bin")
	image.Generate("img.png", shellcode)

}
