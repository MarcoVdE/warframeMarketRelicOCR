package main

import (
	"bufio"
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/vova616/screenshot"
	"image/png"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	screenshotName := "test"
	//screenshotName := "test" + buildFileName()
	//captureScreenshotToAssets(screenshotName)

	imageTest := getTextFromImage("assets/" + screenshotName + ".png")
	//MakeRequest()

	var foundItems []string
	println("Beginning Search.")
	for _, element := range ItemList {
		if strings.Contains(imageTest, element) {
			foundItems = append(foundItems, element)
		}
	}

	fmt.Printf("%+q", foundItems)
}

func getTextFromImage(filename string) string {

	imgEncoded, err := Base64(filename)
	if err != nil {
		log.Fatalln("image could not be encoded, err: " + err.Error())
	}

	message := map[string]interface{}{
		"base64": imgEncoded,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost:8080/base64", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	_ = json.NewDecoder(resp.Body).Decode(&result)

	//return result["data"]
	//log.Println(result)

	//b, err := json.MarshalIndent(result, "", "  ")
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//fmt.Print(string(b))

	//log.Println(result["data"])
	return  result["result"].(string)
}

func Base64(path string) (string, error) {
	imgFile, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := b64.StdEncoding.EncodeToString(buf)

	return imgBase64Str,nil
}

func captureScreenshotToAssets(newFilename string) {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}
	f, err := os.Create("assets/"+ newFilename + ".png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
	f.Close()
}

func buildFileName() string {
	return time.Now().Format("20060102150405")
}