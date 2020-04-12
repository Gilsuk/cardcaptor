package lib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

/*
TokenResp is
*/
type TokenResp struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

/*
TokenReq is
*/
type TokenReq struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

/*
Card is
*/
type Card struct {
	ID            int    `json:"id"`
	Collectible   int    `json:"collectible"`
	Slug          string `json:"slug"`
	ClassID       int    `json:"classId"`
	MultiClassIDs []int  `json:"multiClassIds"`
	MinionTypeID  int    `json:"minionTypeId"`
	CardTypeID    int    `json:"cardTypeId"`
	CardSetID     int    `json:"cardSetId"`
	RarityID      int    `json:"rarityId"`
	ArtistName    string `json:"artistName"`
	Health        int    `json:"health"`
	Attack        int    `json:"attack"`
	ManaCost      int    `json:"manaCost"`
	Name          string `json:"name"`
	Text          string `json:"text"`
	Image         string `json:"image"`
	ImageGold     string `json:"imageGold"`
	FlavorText    string `json:"flavorText"`
	CropImage     string `json:"cropImage"`
}

/*
CardResp is
*/
type CardResp struct {
	Cards     []Card `json:"cards"`
	CardCount int    `json:"cardCount"`
	PageCount int    `json:"pageCount"`
	Page      int    `json:"page"`
}

/*
Export is
*/
func (c *Card) Export() {
	path := "./res/card"
	os.MkdirAll(path, os.ModePerm)

	bytes, _ := json.Marshal(*c)
	ioutil.WriteFile(path+"/"+c.Slug, bytes, os.ModePerm)
}

/*
HasNext is
*/
func (c *CardResp) HasNext() bool {
	if c.Page < c.PageCount {
		return true
	}
	return false
}

/*
Next is
*/
func (c *CardResp) Next() CardResp {
	return CardResp{}
}

/*
RequestToken is
*/
func RequestToken() (respJSON TokenResp) {
	conf := NewConf()
	req := TokenReq{GrantType: "client_credentials", ClientID: conf.ClientID, ClientSecret: conf.ClientSecret}
	reqJSON, _ := json.Marshal(req)
	reqBytes := bytes.NewBuffer(reqJSON)

	resp, err := http.Post("https://apac.battle.net/oauth/token", "application/json", reqBytes)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(respBytes, &respJSON)
	return
}

/*
RequestCards is
*/
func RequestCards(page int) (cardResp CardResp) {
	conf := NewConf()
	base, _ := url.Parse("https://kr.api.blizzard.com")
	base.Path += "hearthstone/cards"
	params := url.Values{}
	params.Add("locale", "ko_KR")
	params.Add("page", strconv.Itoa(page))
	params.Add("access_token", conf.AccessToken)
	base.RawQuery = params.Encode()

	println(base.String())
	resp, err := http.Get(base.String())
	if err != nil {
		println("Request fail")
		return
	}

	defer resp.Body.Close()
	if respBytes, err := ioutil.ReadAll(resp.Body); err == nil {
		json.Unmarshal(respBytes, &cardResp)
	}

	println(cardResp.CardCount)
	println(cardResp.Cards[0].Image)

	for _, card := range cardResp.Cards {
		card.Export()
	}

	return
}
