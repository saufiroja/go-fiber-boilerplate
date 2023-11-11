package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"project/go-fiber-boilerplate/utils/constants"
)

type ISuite interface {
	Result(data *constants.Response, http *http.Response)
	Post(path string, body *bytes.Buffer, data *constants.Response) (*constants.Response, *http.Response, error)
	Get(path string, data *constants.Response) (*constants.Response, *http.Response, error)
	Put(path string, body *bytes.Buffer, data *constants.Response) (*constants.Response, *http.Response, error)
	Delete(path string, data *constants.Response) (*constants.Response, *http.Response, error)
}

type SuiteUtils struct {
	client *http.Client
}

func NewSuiteUtils(client *http.Client) ISuite {
	return &SuiteUtils{
		client: client,
	}
}

func (s *SuiteUtils) sendRequest(method, path string, body *bytes.Buffer) (resp *http.Response, err error) {
	bdy := bytes.NewBuffer(nil)
	if body != nil {
		bdy = body
	}

	req, err := http.NewRequest(method, path, bdy)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err = s.client.Do(req)
	return resp, err
}

func (s *SuiteUtils) Result(data *constants.Response, http *http.Response) {
	decoder := json.NewDecoder(http.Body)
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
}

func (s *SuiteUtils) Post(path string, body *bytes.Buffer, data *constants.Response) (*constants.Response, *http.Response, error) {
	res, err := s.sendRequest("POST", path, body)
	if err != nil {
		return nil, nil, err
	}

	s.Result(data, res)

	return data, res, nil
}

func (s *SuiteUtils) Get(path string, data *constants.Response) (*constants.Response, *http.Response, error) {
	res, err := s.sendRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	s.Result(data, res)

	return data, res, nil
}

func (s *SuiteUtils) Put(path string, body *bytes.Buffer, data *constants.Response) (*constants.Response, *http.Response, error) {
	res, err := s.sendRequest("PUT", path, body)
	if err != nil {
		return nil, nil, err
	}

	s.Result(data, res)

	return data, res, nil
}

func (s *SuiteUtils) Delete(path string, data *constants.Response) (*constants.Response, *http.Response, error) {
	res, err := s.sendRequest("DELETE", path, nil)
	if err != nil {
		return nil, nil, err
	}

	s.Result(data, res)
	return data, res, nil
}
