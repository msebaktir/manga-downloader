package imagehelper

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"os"
)

func readImageConfig(path string) image.Config {
	if reader, err := os.Open(path); err == nil {
		defer reader.Close()
		image, _, err := image.DecodeConfig(reader)
		if err != nil {
			panic("Error while getting image size " + err.Error() + " " + path + " ")
		}
		return image
	}
	panic("Error while getting image size")
}
func readImageFormat(path string) string {
	if reader, err := os.Open(path); err == nil {
		defer reader.Close()
		_, imageType, err := image.DecodeConfig(reader)
		if err != nil {
			panic("Error while getting image size " + err.Error() + " " + path + " ")
		}
		return imageType
	}
	panic("Error while getting image type")
}
func readImage(path string) image.Image {
	if reader, err := os.Open(path); err == nil {
		defer reader.Close()
		image, _, err := image.Decode(reader)
		if err != nil {
			panic("Error while getting image size " + err.Error() + " " + path + " ")
		}
		return image
	}
	panic("Error while getting image")
}
func GetImageSize(path string) (float64, float64) {
	image := readImageConfig(path)
	return float64(image.Width), float64(image.Height)
}
func CalculateSizeForPapaerSize(ratio int, imagePath string) (float64, float64) {
	width, height := GetImageSize(imagePath)
	calculated_width := (width * float64(ratio)) - 20
	calculated_height := (height * float64(ratio)) - 20
	return calculated_width, calculated_height
}
func DecodeJpeg(path string) image.Image {
	// decode jpeg
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("Error while decoding jpeg: " + path)
		panic(err)
	}
	return img
}
func GetImageFormat(path string) string {
	return readImageFormat(path)
}
func ConvertByteToImage(byte []byte) image.Image {
	img, _, err := image.Decode(bytes.NewReader(byte))
	if err != nil {
		panic(err)
	}
	return img
}
func WidthHeightToMM(path string) (float64, float64) {
	width, height := GetImageSize(path)
	return width * 0.2645833333, height * 0.2645833333

}
func SaveImage(image image.Image, path string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	jpeg.Encode(out, image, nil)
	return nil
}
