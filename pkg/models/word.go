package models

type Word struct {
	Origin string `json:"origin"`
}

type TranslateWord struct {
	Origin    string `json:"origin"`
	Translate string `json:"translate"`
}
