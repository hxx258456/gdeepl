/*
Copyright 2022 github.com/hxx258456/gdeepl
*/
package client

import "testing"

func TestGet(t *testing.T) {
	data := &Request{
		Text:       "------------------------------",
		TargetLang: "ZH",
		SourceLang: "auto",
	}
	if err := Post(data); err != nil {
		t.Error(err)
	}
}
