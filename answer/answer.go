package answer

import (
	"go.m3o.com/client"
)

type Answer interface {
	Question(*QuestionRequest) (*QuestionResponse, error)
}

func NewAnswerService(token string) *AnswerService {
	return &AnswerService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type AnswerService struct {
	client *client.Client
}

// Ask a question and receive an instant answer
func (t *AnswerService) Question(request *QuestionRequest) (*QuestionResponse, error) {

	rsp := &QuestionResponse{}
	return rsp, t.client.Call("answer", "Question", request, rsp)

}

type QuestionRequest struct {
	// the question to answer
	Query string `json:"query,omitempty"`
}

type QuestionResponse struct {
	// the answer to your question
	Answer string `json:"answer,omitempty"`
	// any related image
	Image string `json:"image,omitempty"`
	// a related url
	Url string `json:"url,omitempty"`
}
