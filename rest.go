package restgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func Get(url string, resp interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	return do(req, resp)
}

func Post(url string, req, resp interface{}) error {
	d, err := json.Marshal(req)
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(d)
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	return do(request, resp)
}

func do(req *http.Request, resp interface{}) error {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, resp)
	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}
	return err
}
