/*
Copyright 2022 github.com/hxx258456/gdeepl
*/
package client

type Request struct {
	Text       string `json:"text"`
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
}

type Reponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}
