package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

	succeedOrFatal(json.Unmarshal(body, &data))

	os.WriteFile("out/"+filename+".json", body, 0644)

	return err
}
func succeedOrFatal(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func readOrFatal(path string) []byte {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return fileContent
}
func main() {

	inSetsContent := readOrFatal("inSets.json")

	var dataSets []string
	succeedOrFatal(json.Unmarshal(inSetsContent, &dataSets))

	inSetsAmountContent := readOrFatal("inSetAmounts.json")

	var dataSetAmounts []int
	succeedOrFatal(json.Unmarshal(inSetsAmountContent, &dataSetAmounts))

	for i, set := range dataSets {
		GetAndWrite(fmt.Sprintf("https://api.pokemontcg.io/v2/cards?q=set.id:%s&page=1&pageSize=250", set), set, os.Args[1])

		// Some modern sets are bigger than the max allowed pagesize of the backing API
		if dataSetAmounts[i] > 249 {
			GetAndWrite(fmt.Sprintf("https://api.pokemontcg.io/v2/cards?q=set.id:%s&page=2&pageSize=250", set), os.Args[1], set+"_ext")
		}

	}
}
