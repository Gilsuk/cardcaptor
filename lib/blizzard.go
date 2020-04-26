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
	Card          int    `json:"id"`
	Slug          string `json:"slug"`
	Class         int    `json:"classId"`
	Type          int    `json:"cardTypeId"`
	Set           int    `json:"cardSetId"`
	Rarity        int    `json:"rarityId"`
	Race          int    `json:"minionTypeId"`
	Artist        string `json:"artistName"`
	Name          string `json:"name"`
	Text          string `json:"text"`
	Flavor        string `json:"flavorText"`
	Img           string `json:"image"`
	CropImg       string `json:"cropImage"`
	Cost          int    `json:"manaCost"`
	Health        int    `json:"health"`
	Attack        int    `json:"attack"`
	Armor         int    `json:"armor"`
	Collectible   int    `json:"collectible"`
	MultiClassIDs []int  `json:"multiClassIds"`
	Parent        int    `json:"parentId"`
	Child         []int  `json:"childIds"`
	Keyword       []int  `json:"keywordIds"`
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
Meta is
*/
type Meta struct {
	Sets      []Set      `json:"sets"`
	SetGroups []SetGroup `json:"setGroups"`
	Arenas    []int      `json:"arenaIds"`
	Types     []Type     `json:"types"`
	Rarities  []Rarity   `json:"rarities"`
	Classes   []Class    `json:"classes"`
	Races     []Race     `json:"minionTypes"`
	Keywords  []Keyword  `json:"keywords"`
}

/*
Set is
*/
type Set struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	ReleaseDate string `json:"releaseDate"`
	Type        string `json:"type"`
}

/*
SetGroup is
*/
type SetGroup struct {
	Slug     string   `json:"slug"`
	Year     int      `json:"year"`
	Sets     []string `json:"cardSets"`
	Name     string   `json:"name"`
	Standard bool     `json:"standard"`
}

/*
Keyword is
*/
type Keyword struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Ref  string `json:"refText"`
	Text string `json:"text"`
}

/*
Rarity is
*/
type Rarity struct {
	Slug string `json:"slug"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

/*
Class is
*/
type Class struct {
	Slug string `json:"slug"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

/*
Race is
*/
type Race struct {
	Slug string `json:"slug"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

/*
Type is
*/
type Type struct {
	Slug string `json:"slug"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

/*
HSJsonCard is for JSON from https://hearthstonejson.com/
*/
type HSJsonCard struct {
	ID    int      `json:"dbfId"`
	Mechs []string `json:"mechanics"`
	Refs  []string `json:"referencedTags"`
}

// HSJsonResp is
type HSJsonResp []HSJsonCard

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
	println(cardResp.Cards[0].Img)

	for _, card := range cardResp.Cards {
		card.Export()
	}

	return
}
