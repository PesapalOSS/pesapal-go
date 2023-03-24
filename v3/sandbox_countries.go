package pesapalv3

type SandboxCountry int

const (
	Kenya SandboxCountry = iota
	Uganda
	Tanzania
	Malawi
	Rwanda
	Zambia
	Zimbabwe
)

var SandboxCountries = struct {
	Kenya    SandboxCountry
	Uganda   SandboxCountry
	Tanzania SandboxCountry
	Malawi   SandboxCountry
	Rwanda   SandboxCountry
	Zambia   SandboxCountry
	Zimbabwe SandboxCountry
}{
	Kenya:    Kenya,
	Uganda:   Uganda,
	Tanzania: Tanzania,
	Malawi:   Malawi,
	Rwanda:   Rwanda,
	Zambia:   Zambia,
	Zimbabwe: Zimbabwe,
}
