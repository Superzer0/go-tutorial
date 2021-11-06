package app

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Superzer0/go-tutorial/internal/pkg/common"
)

func MergeResults(fileName1 string, fileName2 string, outputFile string) error {
	fmt.Printf("Merging %s and %s into %s", fileName1, fileName2, outputFile)

	wordsFromFirstFile, err := readWords(fileName1)
	if err != nil {
		return fmt.Errorf("error when merging %w", err)
	}

	wordsFromSecondFile, err := readWords(fileName2)
	if err != nil {
		return fmt.Errorf("error when merging %w", err)
	}

	mergedValues := common.MergeWordCounts(wordsFromFirstFile, wordsFromSecondFile)
	sort.Sort(sort.Reverse(mergedValues))

	err = common.SaveOutputFile(mergedValues, outputFile)
	if err != nil {
		return fmt.Errorf("could not merge files %w", err)
	}

	fmt.Printf("Merged files and saved in %s \n", outputFile)
	return nil
}

func readWords(fileName string) (common.PairList, error) {

	file1, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not open file %s %w", fileName, err)
	}
	defer file1.Close()

	return readFileContent(file1)
}

func readFileContent(file *os.File) (common.PairList, error) {

	output := make(common.PairList, 0)
	scanner := bufio.NewScanner(file)

	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		pair, err := parseLine(line, lineNum, file.Name())
		if err != nil {
			return nil, err
		}
		if pair == nil {
			continue
		}

		output = append(output, *pair)
	}

	return output, nil
}

func parseLine(line string, lineNum int, fileName string) (*common.Pair, error) {
	values := strings.Split(line, ":")

	if len(values) < 2 {
		return nil, nil
	}

	if len(values) > 2 {
		return nil, fmt.Errorf("invalid format encountered in %s at %d line", fileName, lineNum)
	}

	if len(values[0]) == 0 {
		return nil, fmt.Errorf("cannot have empty word in %s at %d line", fileName, lineNum)
	}

	occurences, err := strconv.Atoi(values[1])
	if err != nil {
		return nil, fmt.Errorf("could not parse occurences in %s at %d line", fileName, lineNum)
	}

	return &common.Pair{
		Word:       values[0],
		Occurences: occurences,
	}, nil
}
