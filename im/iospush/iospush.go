// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package iospush

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func baseget(posturl string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", posturl, nil)
	req.Header.Add("service", "wukong")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("body: %v", string(body))
}

func Push(apiurl, msg string, apikeys []string) {
	for _, apikey := range apikeys {
		pushurl := fmt.Sprintf("%v/%v/%v", apiurl, apikey,msg)
		go baseget(pushurl)
	}
}