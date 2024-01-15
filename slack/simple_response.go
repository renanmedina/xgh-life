package slack

type SlackSimpleResponse struct {
	ResponseType string `json:"response_type"`
	Body         string `json:"text"`
}

const (
	IN_CHANNEL = "in_channel"
)

func NewSlackSimpleResponse(rtype string, body string) SlackSimpleResponse {
	return SlackSimpleResponse{rtype, body}
}
