package main

import (
	"encoding/json"
	"strings"
	"time"
)

type GeneralError struct {
	Type    string `json:"error_type"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type OurCredentials struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
}

type ApiCredentials struct {
	Token      string        `json:"token"`
	ExpiryDate time.Time     `json:"expiryDate"`
	Error      *GeneralError `json:"error"`
	Status     string        `json:"status"`
	Message    string        `json:"message"`
}

type BillingAddress struct {
	EmailAddress string  `json:"email_address"`
	PhoneNumber  *string `json:"phone_number"`
	CountryCode  *string `json:"country_code"`
	FirstName    *string `json:"first_name"`
	MiddleName   *string `json:"middle_name"`
	LastName     *string `json:"last_name"`
	Line1        *string `json:"line_1"`
	Line2        *string `json:"line_2"`
	City         *string `json:"city"`
	State        *string `json:"state"`
	PostalCode   *string `json:"postal_code"`
	ZipCode      *string `json:"zip_code"`
}

type SubmitOrderReq struct {
	Id              string         `json:"id"`
	Currency        string         `json:"currency"`
	Amount          float64        `json:"amount"`
	Description     string         `json:"description"`
	CallbackUrl     string         `json:"callback_url"`
	CancellationUrl string         `json:"cancellation_url"`
	NotificationId  string         `json:"notification_id"`
	BillingAddress  BillingAddress `json:"billing_address"`
}

type SubmitOrderRes struct {
	OrderTrackingId   string        `json:"order_tracking_id"`
	MerchantReference string        `json:"merchant_reference"`
	RedirectUrl       string        `json:"redirect_url"`
	Error             *GeneralError `json:"error"`
	Status            string        `json:"status"`
}

type TransactionStatusRes struct {
	PaymentMethod            string        `json:"payment_method"`
	Amount                   float64       `json:"amount"`
	CreatedDate              time.Time     `json:"created_date"`
	ConfirmationCode         string        `json:"confirmation_code"`
	PaymentStatusDescription string        `json:"payment_status_description"`
	Description              string        `json:"description"`
	Message                  string        `json:"message"`
	PaymentAccount           string        `json:"payment_account"`
	CallbackUrl              string        `json:"call_back_url"`
	StatusCode               uint          `json:"status_code"`
	MerchantReference        string        `json:"merchant_reference"`
	PaymentStatusCode        string        `json:"payment_status_code"`
	Currency                 string        `json:"currency"`
	Error                    *GeneralError `json:"error"`
	Status                   string        `json:"status"`
}

type IPNRegisterReq struct {
	Url                 string `json:"url"`
	IpnNotificationType string `json:"ipn_notification_type"`
}

type IPNRegisterRes struct {
	Url         string        `json:"url"`
	CreatedDate time.Time     `json:"created_date"`
	IpnId       string        `json:"ipn_id"`
	Error       *GeneralError `json:"error"`
	Status      string        `json:"status"`
}

type IPNGetEndpointsRes []IPNRegisterRes

type IPNNotification struct {
	OrderMerchantReference string `json:"OrderMerchantReference"`
	OrderNotificationType  string `json:"OrderNotificationType"`
	OrderTrackingId        string `json:"OrderTrackingId"`
}

func (r *IPNRegisterRes) UnmarshalJSON(data []byte) error {
	type Alias IPNRegisterRes
	aux := &struct {
		CreatedDate string `json:"created_date"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var formatstamp = "2006-01-02T15:04:05.99"
	if strings.Contains("Z", aux.CreatedDate) {
		formatstamp = "2006-01-02T15:04:05.000Z"
	}
	timestamp, err := time.Parse(formatstamp, aux.CreatedDate)
	if err != nil {
		return err
	}
	r.CreatedDate = timestamp

	return nil
}

func (r *TransactionStatusRes) UnmarshalJSON(data []byte) error {
	type Alias TransactionStatusRes
	aux := &struct {
		CreatedDate string `json:"created_date"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	timestamp, err := time.Parse("2006-01-02T15:04:05.99", aux.CreatedDate)
	if err != nil {
		return err
	}
	r.CreatedDate = timestamp

	return nil
}
