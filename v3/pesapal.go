package pesapalv3

type Pesapal struct {
	client Client // HTTP client used to communicate with the Pesapal API
}

type PesapalConfig struct {
	Sandbox        bool           // True if the configuration is for sandbox environment, otherwise false
	SandboxCountry SandboxCountry // If sandbox is true, the configuration for the sandbox environment
	ConsumerKey    string         // Consumer key for production environment
	ConsumerSecret string         // Consumer secret for production environment
}

func NewClient(c *PesapalConfig) *Pesapal {
	pesapal := &Pesapal{}

	return pesapal
}
