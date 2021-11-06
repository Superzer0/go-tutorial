package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Superzer0/go-tutorial/internal/pkg/common"
)

// Downloads and analyzes files from remote url
// Urls are fetched from settings file
func AnalyzeLinks(file1 string, file2 string) error {
	settings, err := common.ReadSettings()
	if err != nil {
		log.Fatal("Cannot analyze files without settings")
		return err
	}

	fmt.Printf("Read settings: %+v\n", settings)
	downloadWorkChan := make(chan downloadWorkInput, 2)
	analyzeStringChan := make(chan analyzeStringInput, 2)
	doneChan := make(chan bool, 2)
	inputToProcess := []downloadWorkInput{
		{
			Url:      settings.FirstSourceFileUrl,
			FileName: file1,
		},
		{
			Url:      settings.SecondSourceFileUrl,
			FileName: file2,
		},
	}

	startWorkers(downloadWorkChan, analyzeStringChan, doneChan)

	for _, input := range inputToProcess {
		downloadWorkChan <- input
	}

	close(downloadWorkChan)
	defer close(doneChan)
	defer close(analyzeStringChan)

	for i := 0; i < len(inputToProcess); i++ {
		select {
		case <-doneChan:
		case <-time.After(time.Minute * 1):
			log.Fatal("Analyze task timed out after 1 minutes")
			return fmt.Errorf("Analyze task timed out after 1 minute")
		}
	}

	return nil
}

func startWorkers(downloadChan chan downloadWorkInput, analyzChan chan analyzeStringInput, doneChan chan bool) {
	// start 2 workers for download
	for i := 0; i < 2; i++ {
		go downloadWorker(downloadChan, analyzChan)
	}

	// start 1 worker for analyze
	go analyzeWorker(analyzChan, doneChan)
}

func analyzeWorker(inputChan chan analyzeStringInput, doneChan chan bool) {
	for input := range inputChan {

		wordCount := common.RankByWordCount(input.Content)

		out, err := os.Create(input.FileName)
		if err != nil {
			log.Fatalf("Could not save %s. %v", input.FileName, err)
		}

		var itemsToWrite = 50
		if itemsToWrite > len(wordCount) {
			itemsToWrite = len(wordCount)
		}

		for i := 0; i < itemsToWrite; i++ {
			out.WriteString(fmt.Sprintln(wordCount[i].Word))
		}

		out.Close()
		fmt.Printf("Analyzed and saved %s \n", input.FileName)
		doneChan <- true
	}
}

func downloadWorker(inputChan chan downloadWorkInput, outputChan chan analyzeStringInput) {

	for input := range inputChan {
		output, err := common.DownloadFile(input.Url, input.FileName)
		if err != nil {
			log.Fatalf("Could not download and save file %s. Err %v", input.Url, err)
		} else {
			outputChan <- analyzeStringInput{
				Content:  output,
				FileName: input.FileName,
			}
		}
	}

}

type downloadWorkInput struct {
	Url      string
	FileName string
}
type analyzeStringInput struct {
	Content  string
	FileName string
}
