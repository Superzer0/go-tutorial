package common

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	Analyze string = "analyze"
	Merge   string = "merge"
)

func DownloadFile(url string, fileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)

	fmt.Printf("Downloaded %v to %v\n", url, fileName)

	return err
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
