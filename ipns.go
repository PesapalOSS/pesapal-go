package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (p *Pesapal) GetIPNs() (IPNGetEndpointsRes, error) {
	err := p.ensureAuth()
	if err != nil {
		return IPNGetEndpointsRes{}, err
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", p.apiUrl+"/api/URLSetup/GetIpnList", nil)
	if err != nil {
		return IPNGetEndpointsRes{}, err
	}
	req.Header.Add("Authorization", "Bearer "+p.apiCredentials.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return IPNGetEndpointsRes{}, err
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return IPNGetEndpointsRes{}, err
	}
	var data IPNGetEndpointsRes
	if err := json.Unmarshal(resBody, &data); err != nil {
		return IPNGetEndpointsRes{}, err
	}
	return data, nil
}

func (p *Pesapal) RegisterIPN(d IPNRegisterReq) (IPNRegisterRes, error) {
	err := p.ensureAuth()
	if err != nil {
		return IPNRegisterRes{}, err
	}
	jsonBody, _ := json.Marshal(d)
	client := &http.Client{}
	req, err := http.NewRequest("POST", p.apiUrl+"/api/URLSetup/RegisterIPN", bytes.NewBuffer(jsonBody))
	if err != nil {
		return IPNRegisterRes{}, err
	}
	req.Header.Add("Authorization", "Bearer "+p.apiCredentials.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return IPNRegisterRes{}, err
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return IPNRegisterRes{}, err
	}
	var data IPNRegisterRes
	if err := json.Unmarshal(resBody, &data); err != nil {
		return IPNRegisterRes{}, err
	}
	if data.Error != nil {
		return IPNRegisterRes{}, errors.New("Pesapal - " + data.Error.Message)
	}
	return data, nil
}
