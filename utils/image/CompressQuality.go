package image

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func loadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func imgDecode(imgFile io.Reader, quality int) (string, error) {
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %v", err)
	}
	// 创建一个内存缓冲区来存储压缩后的图片数据
	var buf bytes.Buffer
	// 设置JPEG编码选项以降低质量
	opts := &jpeg.Options{Quality: quality} // 质量范围通常是1-100
	// 使用新的质量参数重新编码图片
	err = jpeg.Encode(&buf, img, opts)
	if err != nil {
		return "", fmt.Errorf("failed to encode image: %v", err)
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// CompressQualityAndEncodeToBase64ByFile 压缩JPEG图片质量和转换为Base64编码（不改变尺寸）
// 对于本地图片
func CompressQualityAndEncodeToBase64ByFile(filename string, quality int) (string, error) {
	// 读取原始图片
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open image: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// 将字节流转换为Base64编码
	return imgDecode(file, quality)
}

// CompressQualityAndEncodeToBase64ByUrl 压缩JPEG图片质量和转换为Base64编码（不改变尺寸）
// 对于网络图片
func CompressQualityAndEncodeToBase64ByUrl(url string, quality int) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to get image: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(response.Body)
	tmpFile, err := ioutil.TempFile("", "temp_image_*.jpg") //建立临时文件
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %v", err)
	}
	defer func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name()) //处理完毕删除临时文件
	}()
	_, err = io.Copy(tmpFile, response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to write response body to temp file: %v", err)
	}
	// 读取临时文件内容作为图像
	imgFile, err := os.Open(tmpFile.Name())
	if err != nil {
		return "", fmt.Errorf("failed to open temp file: %v", err)
	}
	defer imgFile.Close()

	// 将字节流转换为Base64编码
	return imgDecode(imgFile, quality)
}
