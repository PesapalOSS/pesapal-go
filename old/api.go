package pesapal

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Pesapal struct {
	apiUrl         string
	apiCredentials ApiCredentials
	ourCredentials OurCredentials
	IPN            IPNRegisterRes
}

type InitParams struct {
	Sandbox           bool
	CustomCredentials bool
	Credentials       *OurCredentials
	IPN               string
}

func Init(ip InitParams) (*Pesapal, error) {
	apiUrl := "https://pay.pesapal.com/v3"
	var ourCredentials OurCredentials = *ip.Credentials
	if ip.Sandbox == true {
		apiUrl = "https://cybqa.pesapal.com/pesapalv3"
		ourCredentials = OurCredentials{ConsumerKey: "qkio1BGGYAXTu2JOfm7XSXNruoZsrqEW", ConsumerSecret: "osGQ364R49cXKeOYSpaOnT++rHs="}
	}
	if ip.CustomCredentials {
		ourCredentials = *ip.Credentials
	}
	apiCredentials, err := auth(apiUrl, ourCredentials)
	if err != nil {
		return nil, err
	}
	p := &Pesapal{
		apiUrl:         apiUrl,
		apiCredentials: apiCredentials,
		ourCredentials: ourCredentials,
		IPN:            IPNRegisterRes{},
	}
	log.Println("PESAPAL: Getting IPNs")
	var ipnRegistered bool = false
	registeredIPNs, err := p.GetIPNs()
	for _, registeredIPN := range registeredIPNs {
		if registeredIPN.Url == ip.IPN {
			log.Printf("PESAPAL: IPN %s Registered", registeredIPN.Url)
			ipnRegistered = true
			p.IPN = registeredIPN
			break
		}
	}
	if !ipnRegistered {
		log.Printf("PESAPAL: Registering IPN: %s", ip.IPN)
		ipn, err := p.RegisterIPN(IPNRegisterReq{Url: ip.IPN, IpnNotificationType: "POST"})
		if err != nil {
			return nil, err
		}
		p.IPN = ipn
	}
	if err != nil {
		return nil, err
	}
	return p, nil
}

func auth(url string, c OurCredentials) (ApiCredentials, error) {
	jsonBody, _ := json.Marshal(c)
	resp, err := http.Post(url+"/api/Auth/RequestToken", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return ApiCredentials{}, err
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ApiCredentials{}, err
	}
	var data ApiCredentials
	if err := json.Unmarshal(resBody, &data); err != nil {
		return ApiCredentials{}, err
	}
	if data.Error != nil {
		return ApiCredentials{}, errors.New("Pesapal - " + data.Error.Message)
	}
	log.Printf("PESAPAL: Authenticated")
	return data, nil
}

func (p *Pesapal) ensureAuth() error {
	currentTime := time.Now()

	if currentTime.Before(p.apiCredentials.ExpiryDate) {
		return nil
	}

	credentials, err := auth(p.apiUrl, p.ourCredentials)
	if err != nil {
		return err
	}
	p.apiCredentials = credentials
	return nil
}
