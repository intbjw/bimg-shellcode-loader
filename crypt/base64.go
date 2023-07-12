package crypt

import (
	"encoding/base64"
)

// 自定义base64编码表
const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// 自定义base64编码
func CustomBase64Encode(data []byte) string {
	// 创建一个自定义base64编码器
	encoder := base64.NewEncoding(ShuffleString(base64Table))

	// 编码数据
	encoded := encoder.EncodeToString(data)

	return encoded
}

// 自定义base64解码
func CustomBase64Decode(encoded string) ([]byte, error) {
	// 创建一个自定义base64解码器
	decoder := base64.NewEncoding(ShuffleString(base64Table))

	// 解码数据
	decoded, err := decoder.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

func ShuffleString(s string) string {
	// 将字符串转换为rune数组
	runes := []rune(s)

	// 打乱rune数组顺序
	for i := len(runes) - 1; i > 0; i-- {
		j := i - 1
		if j < 0 {
			j = 0
		}
		runes[i], runes[j] = runes[j], runes[i]
	}

	// 将rune数组转换为字符串
	return string(runes)
}
