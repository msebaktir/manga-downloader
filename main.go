package main

import (
	"fmt"

	"eminbaktir.com/manga-downloader/chapter"
	chapterdownloader "eminbaktir.com/manga-downloader/chapter_downloader"
	"eminbaktir.com/manga-downloader/constants"
	fileService "eminbaktir.com/manga-downloader/fileservice"
	imagehelper "eminbaktir.com/manga-downloader/image_helper"
	pdfhelper "eminbaktir.com/manga-downloader/pdf_helper"
)

var BASE_URL string = "https://tortuga-ceviri.com/manga/berserk/"
var CHAPTERS_AJAX_URL string = "https://tortuga-ceviri.com/manga/berserk/ajax/chapters/"
var Priorty int = 0
var selectedPaperSize constants.PaperSize = constants.A4

func DownloadChapters() {
	chapters := chapter.GetChapters(CHAPTERS_AJAX_URL)
	fmt.Println("Founded chapters: " + fmt.Sprint(len(chapters)))
	fileService.PrepareFolders()
	index := 0
	for _, chapter := range chapters {
		fmt.Println("Chapter: " + chapter.Title + " Downloading... " + fmt.Sprint(index) + "/" + fmt.Sprint(len(chapters)))
		if !fileService.IsChapterExist(chapter.Title) {
			fmt.Println("\tCreating chapter folder: " + chapter.Title)
			fileService.CreateChapterFolder(chapter.Title)
			fmt.Println("\tDownloading chapter: " + chapter.Title)
			chapterdownloader.DonwloadChapter(chapter)
		}
		index++
	}
}
func CalculateImagePerPage(path string) (int, int) {
	width, height := imagehelper.GetImageSize(path)
	width_ratio := float32(selectedPaperSize.WidthInPixel) / float32(width)
	calculated_height := float32(height) * width_ratio
	count_of_image_in_a_page := float32(selectedPaperSize.HeightInPixel) / calculated_height
	floor_height := int(count_of_image_in_a_page)
	return floor_height, int(width_ratio)

}

func convertToPdf() {
	fmt.Println("Converting to pdf...")
	pages := fileService.ReadAllChapters()
	for _, page := range pages {
		fmt.Println("\tConverting chapter: " + page.Title)
		imagePerAPage, ratio := CalculateImagePerPage(page.Images[0].Path)
		fmt.Println("\t\tImage per page: " + fmt.Sprint(imagePerAPage) + " Ratio: " + fmt.Sprint(ratio))
		pdfhelper.CreatePdfWithAllImages(page, imagePerAPage, ratio, selectedPaperSize)
		// err := pdf.OutputFileAndClose("pdf/" + page.Title + ".pdf")
		// if err != nil {
		// 	panic("Error while creating pdf: " + err.Error())
		// }
		// fmt.Println("\t\tPdf created: " + page.Title + ".pdf")
	}
}

func main() {
	selectedPaperSize = constants.A4
	DownloadChapters()
	convertToPdf()
}
