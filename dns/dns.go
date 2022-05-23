package dns

import (
	"go.m3o.com/client"
)

type Dns interface {
	Query(*QueryRequest) (*QueryResponse, error)
}

func NewDnsService(token string) *DnsService {
	return &DnsService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type DnsService struct {
	client *client.Client
}

// Query an addresss
func (t *DnsService) Query(request *QueryRequest) (*QueryResponse, error) {

	rsp := &QueryResponse{}
	return rsp, t.client.Call("dns", "Query", request, rsp)

}

type Answer struct {
	// time to live
	Ttl int32 `json:"TTL,omitempty"`
	// the answer
	Data string `json:"data,omitempty"`
	// name resolved
	Name string `json:"name,omitempty"`
	// type of record
	Type int32 `json:"type,omitempty"`
}

type QueryRequest struct {
	// type of query e.g A, AAAA, MX, SRV
	Type string `json:"type,omitempty"`
	// name to resolve
	Name string `json:"name,omitempty"`
}

type QueryResponse struct {
	Ad       bool       `json:"AD,omitempty"`
	Ra       bool       `json:"RA,omitempty"`
	Provider string     `json:"provider,omitempty"`
	Question []Question `json:"question,omitempty"`
	Status   int32      `json:"status,omitempty"`
	Cd       bool       `json:"CD,omitempty"`
	Rd       bool       `json:"RD,omitempty"`
	Tc       bool       `json:"TC,omitempty"`
	Answer   []Answer   `json:"answer,omitempty"`
}

type Question struct {
	// name to query
	Name string `json:"name,omitempty"`
	// type of record
	Type int32 `json:"type,omitempty"`
}
