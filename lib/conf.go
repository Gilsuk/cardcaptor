package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

/*
Conf is
*/
type Conf struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
}

/*
NewConf is parsed values from conf.json
*/
func NewConf() (conf Conf) {
	jsonFile, err := os.Open("./conf.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &conf)

	return
}
