package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var client *http.Client

func GetAndWrite(url string, filename string, apiKey string) error {
	client = &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("X-Api-Key", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var data map[string]interface{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	errr := json.Unmarshal(body, &data)
	if errr != nil {
		return err
	}

	os.WriteFile("out/"+filename+".json", body, 0644)
	return err
}
func main() {
	fileContent, err := os.ReadFile("inSets.json")
	if err != nil {
	}
	var dataSets []string
	err = json.Unmarshal(fileContent, &dataSets)
	if err != nil {
	}
	inSetAmounts, err := os.ReadFile("inSetAmounts.json")
	if err != nil {
	}
	var dataSetAmounts []int
	err = json.Unmarshal(inSetAmounts, &dataSetAmounts)
	if err != nil {
	}
	for i, set := range dataSets {
		GetAndWrite(fmt.Sprintf("https://api.pokemontcg.io/v2/cards?q=set.id:%s&page=1&pageSize=250", set), set, os.Args[1])
		if dataSetAmounts[i] > 250 {
			GetAndWrite(fmt.Sprintf("https://api.pokemontcg.io/v2/cards?q=set.id:%s&page=2&pageSize=250", set), os.Args[1], set+"_ext")
		}

	}
}
