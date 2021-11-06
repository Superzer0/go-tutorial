package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
)

const (
	Analyze string = "analyze"
	Merge   string = "merge"
)

func DownloadFile(url string, fileName string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	fmt.Printf("Downloaded %v \n", url)
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return string(bodyBytes), nil
	}
	return "", fmt.Errorf("not successful http response status %d", resp.StatusCode)
}

type AppSettings struct {
	FirstSourceFileUrl  string `json:"firstSourceFile"`
	SecondSourceFileUrl string `json:"secondSourceFile"`
}

func ReadSettings() (*AppSettings, error) {
	jsonFile, err := os.Open("settings.json")
	if err != nil {
		return nil, fmt.Errorf("cannot read settings.json! %v ", err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var appSettings AppSettings
	json.Unmarshal(byteValue, &appSettings)
	return &appSettings, nil
}

func RankByWordCount(text string) PairList {
	wordMap := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		wordMap[strings.ToLower(word)] += 1
	}

	pl := NewPairList(wordMap)
	sort.Sort(sort.Reverse(pl))
	return pl
}

func MergeWordCounts(a PairList, b PairList) PairList {
	wordMap := make(map[string]int, len(a))

	for _, pair := range append(a, b...) {
		wordMap[pair.Word] += pair.Occurences
	}
	return NewPairList(wordMap)
}

func NewPairList(dic map[string]int) PairList {
	pl := make(PairList, len(dic))
	var i = 0
	for k, v := range dic {
		pl[i] = Pair{Word: k, Occurences: v}
		i++
	}
	return pl
}

type Pair struct {
	Word       string
	Occurences int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Occurences < p[j].Occurences }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func SaveOutputFile(wordsRank PairList, fileName string) error {
	out, err := os.Create(fileName)

	if err != nil {
		return fmt.Errorf("could not save %s. %w", fileName, err)
	}
	defer out.Close()

	for i := 0; i < len(wordsRank); i++ {
		_, err := out.WriteString(fmt.Sprintf("%s:%d\n", wordsRank[i].Word, wordsRank[i].Occurences))
		if err != nil {
			return err
		}
	}

	return nil
}
