package crypt

import (
	"bufio"
	"bytes"
	"github.com/auyer/steganography"
	"image/png"
	"io"
	"log"
	"os"
)

func StegEncode(imgFile string, msg []byte) {

	inFile, _ := os.Open(imgFile)
	reader := bufio.NewReader(inFile)
	img, _ := png.Decode(reader)
	w := new(bytes.Buffer)                   // buffer that will recieve the results
	err := steganography.Encode(w, img, msg) // Encode the message into the image
	if err != nil {
		log.Printf("Error Encoding file %v", err)
		return
	}
	outFile, _ := os.Create("out_file.png") // create file
	w.WriteTo(outFile)                      // write buffer to it
	outFile.Close()
}

func StegDecode(imgReader io.Reader) []byte {
	reader := bufio.NewReader(imgReader)
	img, _ := png.Decode(reader)
	sizeOfMessage := steganography.GetMessageSizeFromImage(img)
	msg := steganography.Decode(sizeOfMessage, img)
	return msg
}
