package lib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
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
	Arenas    []Arena    `json:"arenaIds"`
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
Type is
*/
type Type struct {
	Slug string `json:"slug"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Race is
type Race Type

// Class is
type Class Type

// Rarity is
type Rarity Type

// Arena is
type Arena int

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
func (c *CardResp) Next(accessToken string) (CardResp, error) {
	return RequestCards(c.Page+1, accessToken)
}

/*
RequestCards is
*/
func RequestCards(page int, accessToken string) (cardResp CardResp, err error) {
	base, _ := url.Parse("https://kr.api.blizzard.com")
	base.Path += "hearthstone/cards"
	params := url.Values{}
	params.Add("locale", "ko_KR")
	params.Add("page", strconv.Itoa(page))
	params.Add("collectible", "0,1")
	params.Add("access_token", accessToken)
	base.RawQuery = params.Encode()

	resp, err := http.Get(base.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if respBytes, err := ioutil.ReadAll(resp.Body); err == nil {
		json.Unmarshal(respBytes, &cardResp)
	}

	return
}

// CrawlMetadata is
func CrawlMetadata(accessToken string) (meta Meta, err error) {
	base, _ := url.Parse("https://kr.api.blizzard.com")
	base.Path += "hearthstone/metadata"
	params := url.Values{}
	params.Add("locale", "ko_KR")
	params.Add("access_token", accessToken)
	base.RawQuery = params.Encode()

	resp, err := http.Get(base.String())
	if err != nil {
		return meta, err
	}

	defer resp.Body.Close()
	if respBytes, err := ioutil.ReadAll(resp.Body); err == nil {
		json.Unmarshal(respBytes, &meta)
	}

	return
}
