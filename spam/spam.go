package spam

import (
	"go.m3o.com/client"
)

type Spam interface {
	Classify(*ClassifyRequest) (*ClassifyResponse, error)
}

func NewSpamService(token string) *SpamService {
	return &SpamService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type SpamService struct {
	client *client.Client
}

// Check whether an email is likely to be spam based on its attributes
func (t *SpamService) Classify(request *ClassifyRequest) (*ClassifyResponse, error) {

	rsp := &ClassifyResponse{}
	return rsp, t.client.Call("spam", "Classify", request, rsp)

}

type ClassifyRequest struct {
	// The raw body of the email including headers etc per RFC 822. Alternatively, use the other parameters to correctly format the message
	EmailBody string `json:"email_body,omitempty"`
	// The email address it has been sent from
	From string `json:"from,omitempty"`
	// the HTML version of the email body
	HtmlBody string `json:"html_body,omitempty"`
	// The subject of the email
	Subject string `json:"subject,omitempty"`
	// the plain text version of the email body
	TextBody string `json:"text_body,omitempty"`
	// The email address it is being sent to
	To string `json:"to,omitempty"`
}

type ClassifyResponse struct {
	// The rules that have contributed to this score
	Details []string `json:"details,omitempty"`
	// Is it spam? Returns true if its score is > 5
	IsSpam bool `json:"is_spam,omitempty"`
	// The score evaluated for this email. A higher number means it is more likely to be spam
	Score float64 `json:"score,omitempty"`
}
