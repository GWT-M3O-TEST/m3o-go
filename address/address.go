package address

import (
	"go.m3o.com/client"
)

type Address interface {
	LookupPostcode(*LookupPostcodeRequest) (*LookupPostcodeResponse, error)
}

func NewAddressService(token string) *AddressService {
	return &AddressService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type AddressService struct {
	client *client.Client
}

// Lookup a list of UK addresses by postcode
func (t *AddressService) LookupPostcode(request *LookupPostcodeRequest) (*LookupPostcodeResponse, error) {

	rsp := &LookupPostcodeResponse{}
	return rsp, t.client.Call("address", "LookupPostcode", request, rsp)

}

type LookupPostcodeRequest struct {
	// UK postcode e.g SW1A 2AA
	Postcode string `json:"postcode,omitempty"`
}

type LookupPostcodeResponse struct {
	Addresses []Record `json:"addresses,omitempty"`
}

type Record struct {
	// dependent locality
	Locality string `json:"locality,omitempty"`
	// the postcode
	Postcode string `json:"postcode,omitempty"`
	// street name
	Street string `json:"street,omitempty"`
	// building name
	BuildingName string `json:"building_name,omitempty"`
	// line two of address
	LineTwo string `json:"line_two,omitempty"`
	// organisation if present
	Organisation string `json:"organisation,omitempty"`
	// the premise
	Premise string `json:"premise,omitempty"`
	// the complete address
	Summary string `json:"summary,omitempty"`
	// post town
	Town string `json:"town,omitempty"`
	// the county
	County string `json:"county,omitempty"`
	// line one of address
	LineOne string `json:"line_one,omitempty"`
}
