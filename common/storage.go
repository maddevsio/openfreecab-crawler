package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/maddevsio/openfreecab-crawler/service/data"
)

func SaveDriver(storageUrl string, driver data.StorageDriver) error {
	client := &http.Client{}
	payload, err := json.Marshal(driver)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/add/", storageUrl)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
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
