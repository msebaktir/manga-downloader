package fileService

import "os"

var BASE_FOLDER string = "BERSERK"

type PageImage struct {
	Path  string
	Title string
}

type ChapterFolder struct {
	Title  string
	Images []PageImage
}

func PrepareFolders() {
	if !isBaseFolderExist() {
		createBaseFolder()
	}
}

func isBaseFolderExist() bool {
	if _, err := os.Stat(BASE_FOLDER); err == nil {
		return true
	}
	return false
}
func createBaseFolder() bool {
	if err := os.Mkdir(BASE_FOLDER, os.ModePerm); err != nil {
		return false
	}
	return true
}

func IsChapterExist(chapterName string) bool {
	if _, err := os.Stat(BASE_FOLDER + "/" + chapterName); err == nil {
		return true
	}
	return false
}
func GetFolderPath(chapterName string) string {
	return BASE_FOLDER + "/" + chapterName
}

func CreateChapterFolder(chapterName string) bool {
	if err := os.Mkdir(BASE_FOLDER+"/"+chapterName, os.ModePerm); err != nil {
		return false
	}
	return true
}

func readAllImages(path string) []PageImage {
	var images []PageImage
	combinedPath := BASE_FOLDER + "/" + path
	files, _ := os.ReadDir(combinedPath)
	for _, file := range files {
		if file.Name() == ".DS_Store" {
			continue
		}
		images = append(images, PageImage{Path: combinedPath + "/" + file.Name(), Title: file.Name()})
	}
	return images

}

func ReadAllChapters() []ChapterFolder {
	var pages []ChapterFolder
	files, _ := os.ReadDir(BASE_FOLDER)
	for _, file := range files {
		if file.Name() == ".DS_Store" {
			continue
		}
		pages = append(pages, ChapterFolder{Title: file.Name(), Images: readAllImages(file.Name())})
	}
	return pages
}
