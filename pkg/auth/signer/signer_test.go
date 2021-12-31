package signer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var s = Signer{
	ID:     "liaoxw@102er.com",
	Secret: "猜猜看",
}

func TestSigner_Sing(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://localhost/api/v1/user", ioutil.NopCloser(bytes.NewBuffer([]byte(""))))
	r.Header.Add("content-type", "application/json")
	err := s.Sing(r)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(r.Header)
	client := http.DefaultClient
	resp, err := client.Do(r)
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.Request)
	t.Log(resp)
	t.Log(string(body))
}
