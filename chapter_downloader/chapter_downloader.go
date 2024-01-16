package chapterdownloader

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"eminbaktir.com/manga-downloader/chapter"
	fileService "eminbaktir.com/manga-downloader/fileservice"
	imagehelper "eminbaktir.com/manga-downloader/image_helper"
	"golang.org/x/net/html"
)

type PageImage struct {
	Url    string
	Number string
}

func findImages(n *html.Node) []PageImage {
	var images []PageImage

	if n.Type == html.ElementNode && n.Data == "img" {
		var tmpImage PageImage
		for _, a := range n.Attr {
			if a.Key == "src" {
				tmpImage.Url = a.Val
			}
			if a.Key == "id" {
				tmpImage.Number = a.Val
			}
		}
		tmpImage.Url = strings.ReplaceAll(tmpImage.Url, "\t", "")
		tmpImage.Url = strings.ReplaceAll(tmpImage.Url, "\n", "")
		images = append(images, tmpImage)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		images = append(images, findImages(c)...)
	}
	return images
}

func findContent(n *html.Node) *html.Node {
	// <div class="reading-content"></div>
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "reading-content" {
				return n
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := findContent(c); result != nil {
			return result
		}
	}
	return nil

}

func downloadImage(image PageImage, path string) {
	res, err := http.Get(image.Url)
	if err != nil {
		fmt.Println("Error while downloading image: " + image.Number)
	}
	defer res.Body.Close()
	saveImage(res.Body, image.Number, path)

}

func saveImage(image io.ReadCloser, number string, path string) {
	// file, err := os.Create(path + "/" + number + ".jpg")
	// if err != nil {
	// 	fmt.Println("Error while creating image: " + number)
	// 	return
	// }
	// defer file.Close()
	// io.Copy(file, image)
	body, err := io.ReadAll(image)
	if err != nil {
		fmt.Println("Error while reading image: " + number)
		return
	}
	// err = os.WriteFile(path+"/"+number+".jpg", body, 0644)
	converted_image := imagehelper.ConvertByteToImage(body)
	err = imagehelper.SaveImage(converted_image, path+"/"+number+".jpg")
	if err != nil {
		fmt.Println("Error while saving image: " + number)
		return
	}

}

func saveImages(images []PageImage, path string) {
	for _, image := range images {
		downloadImage(image, path)
		fmt.Println("\t\tDownloaded image: " + image.Number)
	}
}

func DonwloadChapter(chapter chapter.Chapter) {
	res, err := http.Get(chapter.Url)
	if err != nil {
		fmt.Println("Error while downloading chapter: " + chapter.Title)
		return
	}
	defer res.Body.Close()
	doc, err := html.Parse(res.Body)
	if err != nil {
		fmt.Println("Error while parsing chapter: " + chapter.Title)
		return
	}
	content := findContent(doc)
	images := findImages(content)
	fmt.Println("\t\tFounded images: " + fmt.Sprint(len(images)))
	saveImages(images, fileService.GetFolderPath(chapter.Title))

}
