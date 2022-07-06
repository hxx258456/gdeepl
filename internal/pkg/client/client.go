/*
Copyright 2022 github.com/hxx258456/gdeepl
*/

// post requester
//
package client

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/hxx258456/gdeepl/internal/pkg/utils"
	jsoniter "github.com/json-iterator/go"
)

func Post(request *Request) error {

	// gen body
	req, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(request)
	if err != nil {
		return err
	}
	body := strings.NewReader(utils.B2s(req))

	// request
	resp, err := http.Post("http://121.196.195.20:8080/translate", "application/json", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// read data
	if resp.StatusCode != http.StatusOK {
		return errors.New("request failure: " + resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	result := &Reponse{}
	if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, result); err != nil {
		return err
	}
	if result.Code != 200 {
		return errors.New("request failure: " + utils.B2s(data))
	}
	log.Println(result.Data)

	return nil
}
