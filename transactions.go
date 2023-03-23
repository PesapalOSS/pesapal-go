package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (p *Pesapal) CreateTransaction(d SubmitOrderReq) (SubmitOrderRes, error) {
	err := p.ensureAuth()
	if err != nil {
		return SubmitOrderRes{}, err
	}
	jsonBody, _ := json.Marshal(d)
	client := &http.Client{}
	req, err := http.NewRequest("POST", p.apiUrl+"/api/Transactions/SubmitOrderRequest", bytes.NewBuffer(jsonBody))
	if err != nil {
		return SubmitOrderRes{}, err
	}
	req.Header.Add("Authorization", "Bearer "+p.apiCredentials.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return SubmitOrderRes{}, err
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return SubmitOrderRes{}, err
	}
	var data SubmitOrderRes
	if err := json.Unmarshal(resBody, &data); err != nil {
		return SubmitOrderRes{}, err
	}
	if data.Error != nil {
		return SubmitOrderRes{}, errors.New("Pesapal - " + data.Error.Message)
	}
	return data, nil
}

func (p *Pesapal) GetTransactionStatus(id string) (TransactionStatusRes, error) {
	err := p.ensureAuth()
	if err != nil {
		return TransactionStatusRes{}, err
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", p.apiUrl+"/api/Transactions/GetTransactionStatus?orderTrackingId="+id, nil)
	if err != nil {
		return TransactionStatusRes{}, err
	}
	req.Header.Add("Authorization", "Bearer "+p.apiCredentials.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return TransactionStatusRes{}, err
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return TransactionStatusRes{}, err
	}
	var data TransactionStatusRes
	if err := json.Unmarshal(resBody, &data); err != nil {
		return TransactionStatusRes{}, err
	}
	// if data.Error != nil {
	// 	return TransactionStatusRes{}, errors.New("Pesapal - " + data.Error.Message)
	// }
	return data, nil
}
