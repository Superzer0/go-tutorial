package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		log.Fatalf("Cannot read settings.json! %v \n", err)
		return nil, err
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
		wordMap[word] += 1
	}

	pl := make(PairList, len(wordMap))
	var i = 0
	for k, v := range wordMap {
		pl[i] = Pair{Word: k, Occurences: v}
		i++
	}

	sort.Sort(sort.Reverse(pl))
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
