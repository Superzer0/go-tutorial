package app

import (
	"fmt"
	"log"

	"github.com/Superzer0/go-tutorial/internal/pkg/common"
)

// Downloads and analyzes files from remote url
// Urls are fetched from settings file
func AnalyzeLinks(file1 string, file2 string) {
	settings, err := common.ReadSettings()
	if err != nil {
		log.Fatal("Cannot analyze files without settings")
		return
	}

	fmt.Printf("Read settings: %+v\n", settings)

	// Add error handling
	common.DownloadFile(settings.FirstSourceFileUrl, file1)

	common.DownloadFile(settings.SecondSourceFileUrl, file2)

}
