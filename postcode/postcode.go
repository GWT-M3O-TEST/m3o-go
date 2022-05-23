package postcode

import (
	"go.m3o.com/client"
)

type Postcode interface {
	Lookup(*LookupRequest) (*LookupResponse, error)
	Random(*RandomRequest) (*RandomResponse, error)
	Validate(*ValidateRequest) (*ValidateResponse, error)
}

func NewPostcodeService(token string) *PostcodeService {
	return &PostcodeService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type PostcodeService struct {
	client *client.Client
}

// Lookup a postcode to retrieve the related region, county, etc
func (t *PostcodeService) Lookup(request *LookupRequest) (*LookupResponse, error) {

	rsp := &LookupResponse{}
	return rsp, t.client.Call("postcode", "Lookup", request, rsp)

}

// Return a random postcode and its related info
func (t *PostcodeService) Random(request *RandomRequest) (*RandomResponse, error) {

	rsp := &RandomResponse{}
	return rsp, t.client.Call("postcode", "Random", request, rsp)

}

// Validate a postcode.
func (t *PostcodeService) Validate(request *ValidateRequest) (*ValidateResponse, error) {

	rsp := &ValidateResponse{}
	return rsp, t.client.Call("postcode", "Validate", request, rsp)

}

type LookupRequest struct {
	// UK postcode e.g SW1A 2AA
	Postcode string `json:"postcode,omitempty"`
}

type LookupResponse struct {
	// e.g 51.50354
	Latitude float64 `json:"latitude,omitempty"`
	// e.g -0.127695
	Longitude float64 `json:"longitude,omitempty"`
	// UK postcode e.g SW1A 2AA
	Postcode string `json:"postcode,omitempty"`
	// related region e.g London
	Region string `json:"region,omitempty"`
	// e.g St James's
	Ward string `json:"ward,omitempty"`
	// country e.g United Kingdom
	Country string `json:"country,omitempty"`
	// e.g Westminster
	District string `json:"district,omitempty"`
}

type RandomRequest struct {
}

type RandomResponse struct {
	// country e.g United Kingdom
	Country string `json:"country,omitempty"`
	// e.g Westminster
	District string `json:"district,omitempty"`
	// e.g 51.50354
	Latitude float64 `json:"latitude,omitempty"`
	// e.g -0.127695
	Longitude float64 `json:"longitude,omitempty"`
	// UK postcode e.g SW1A 2AA
	Postcode string `json:"postcode,omitempty"`
	// related region e.g London
	Region string `json:"region,omitempty"`
	// e.g St James's
	Ward string `json:"ward,omitempty"`
}

type ValidateRequest struct {
	// UK postcode e.g SW1A 2AA
	Postcode string `json:"postcode,omitempty"`
}

type ValidateResponse struct {
	// Is the postcode valid (true) or not (false)
	Valid bool `json:"valid,omitempty"`
}
