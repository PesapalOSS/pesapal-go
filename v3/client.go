package pesapalv3

import "net/http"

const (
	prodURL    = "https://pay.pesapal.com/v3"
	sandboxURL = "https://cybqa.pesapal.com/pesapalv3"
)

type Client struct {
	client *http.Client
}
