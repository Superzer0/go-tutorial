package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Superzer0/go-tutorial/internal/pkg/common"
)

type cmdInput struct {
	Mode           string
	FirstFileName  string
	SecondFileName string
	OutputFileName string
}

func main() {
	if cmdParams, error := parseInput(); error == nil {
		fmt.Println(cmdParams)
	} else {
		os.Exit(1)
	}
}

func parseInput() (*cmdInput, error) {

	mode := flag.String("mode", common.Analyze, "Program mode. Use analyze or merge")
	firstFileName := flag.String("file1", "out1.txt", "First file name")
	secondFileName := flag.String("file2", "out2.txt", "Second file name")
	outputFileName := flag.String("outFile", "out3.txt", "Output file name")

	flag.Parse()

	cmdInputParams := cmdInput{
		Mode:           *mode,
		FirstFileName:  *firstFileName,
		SecondFileName: *secondFileName,
		OutputFileName: *outputFileName,
	}

	if cmdInputParams.Mode != common.Analyze &&
		cmdInputParams.Mode != common.Merge {
		log.Fatalf("Unrecognized option %s \n", cmdInputParams.Mode)
		return nil, errors.New("Unrecognized option")
	}

	log.Printf("Running with: %+v \n", cmdInputParams)

	return &cmdInputParams, nil
}
