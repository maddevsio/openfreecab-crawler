package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/maddevsio/openfreecab-crawler/service/data"
)

func request(url string, reader io.Reader) error {

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	if !response.Success {
		return errors.New(response.Message)
	}
	return nil
}

func SaveDriver(storageUrl string, driver data.StorageDriver) error {
	payload, err := json.Marshal(driver)
	if err != nil {
		return err
	}
	return request(fmt.Sprintf("%s/add/", storageUrl), bytes.NewBuffer(payload))
}

func CleanStorage(storageUrl, companyName string) error {
	return request(fmt.Sprintf("%s/clean/%s/", storageUrl, companyName), nil)
}
