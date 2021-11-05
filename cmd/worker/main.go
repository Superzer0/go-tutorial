package main

import (
	"flag"
	"log"
	"os"

	"github.com/Superzer0/go-tutorial/internal/app"
	"github.com/Superzer0/go-tutorial/internal/pkg/common"
)

type cmdInput struct {
	Mode           string
	FirstFileName  string
	SecondFileName string
	OutputFileName string
}

func main() {

	cmdParams := parseInput()

	switch cmdParams.Mode {
	case common.Analyze:
		app.AnalyzeLinks(cmdParams.FirstFileName, cmdParams.SecondFileName)
	case common.Merge:
		app.MergeResults(cmdParams.FirstFileName, cmdParams.SecondFileName, cmdParams.OutputFileName)
	default:
		log.Fatalf("Unrecognized option %s \n", cmdParams.Mode)
		os.Exit(1)
	}
}

func parseInput() *cmdInput {

	mode := flag.String("mode", common.Analyze, "Program mode. Use analyze or merge")
	firstFileName := flag.String("file1", "out1.txt", "First downloaded file name")
	secondFileName := flag.String("file2", "out2.txt", "Second downloaded file name")
	outputFileName := flag.String("mergeFile", "out3.txt", "Merge output file name")

	flag.Parse()

	cmdInputParams := cmdInput{
		Mode:           *mode,
		FirstFileName:  *firstFileName,
		SecondFileName: *secondFileName,
		OutputFileName: *outputFileName,
	}
	log.Printf("Running with: %+v \n", cmdInputParams)

	return &cmdInputParams
}
