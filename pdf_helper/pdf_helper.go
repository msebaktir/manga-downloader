package pdfhelper

import (
	"fmt"

	"eminbaktir.com/manga-downloader/constants"
	fileService "eminbaktir.com/manga-downloader/fileservice"
	"github.com/jung-kurt/gofpdf"
)

var RATIO int = 0

func calculateBottomFromPageHeight(height float64, paper constants.PaperSize) float64 {
	return paper.HeightInPixel - height - 20
}

func addImageToTopOfPage(image fileService.PageImage, pdf *gofpdf.Fpdf) {
	pdf.ImageOptions(image.Path, 10, 10, 190, 0, true, gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true}, 0, "")
}
func addImageToBottomOfPage(image fileService.PageImage, pdf *gofpdf.Fpdf) {

	pdf.ImageOptions(image.Path, 10, 297/2, 190, 0, true, gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true}, 0, "")
}

func CreatePdfWithAllImages(book fileService.ChapterFolder, countOfImage int, ratio int, paperSize constants.PaperSize) {

	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr
	RATIO = ratio
	pdf := gofpdf.New("P", "mm", paperSize.Name, "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 46)
	pdf.Cell(10, 250, book.Title+" BOLUM")
	// pdf.AddPage()
	// index := 0
	for _, image := range book.Images {
		// imagehelper.DecodeJpeg(image.Path)
		addImageToTopOfPage(image, pdf)
	}
	totalPage := pdf.PageNo()
	fmt.Println("Total page: " + fmt.Sprint(totalPage) + " Count of image: " + fmt.Sprint(len(book.Images)))
	err := pdf.OutputFileAndClose("pdf/" + book.Title + ".pdf")
	if err != nil {
		panic("Error while creating pdf: " + err.Error())
	}

}
