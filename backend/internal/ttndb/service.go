package ttndb

// a service for retrieval of messages from the Data Store integration of The Things Network

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"vineguard/environ"
)

type Svc interface {
	GetDevices() ([]string, error)
	GetAllFromDevice(string) (LSB50Msgs, error)
}

type svc struct {
	env *environ.Environ
}

func NewService() Svc {
	env, err := environ.Get()
	if err != nil {
		logrus.Fatalf("could not import environment variables. %s", err)
	}
	return &svc{env: env}
}

func (s *svc) queryTTN(urlSuffix string, outputMsgs interface{}) error {
	const (
		baseUrl = "https://vineguard-lsn50.data.thethingsnetwork.org/api/v2"
	)

	url := baseUrl + urlSuffix
	c := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("could not create http client. %s", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("key %s", s.env.TTNKey))
	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("problem querying TTN DB. %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("could not parse message body. %s", err)
		}

		err = json.Unmarshal(body, &outputMsgs)
		if err != nil {
			return fmt.Errorf("could not parse json. %s", err)
		}
		return nil
	} else {
		return fmt.Errorf("received non-200 code from TTN DB API. %d", resp.StatusCode)
	}
}

func (s *svc) GetDevices() ([]string, error) {
	// retrieve all devices whose data is accessible at this endpoint
	const getDeviceSuffix = "/devices"
	var outputMsgs []string
	err := s.queryTTN(getDeviceSuffix, &outputMsgs)
	return outputMsgs, err
}

func (s *svc) GetAllFromDevice(deviceId string) (LSB50Msgs, error) {
	// retrieve all of the data from a given device

	var outputMsgs LSB50Msgs
	getAllDataSuffix := fmt.Sprintf("/query/%s?last=7d", deviceId)
	err := s.queryTTN(getAllDataSuffix, &outputMsgs)
	return outputMsgs, err
}
