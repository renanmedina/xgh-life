package slack

type SlackSimpleResponse struct {
	ResponseType string `json:"response_type"`
	Body         string `json:"text"`
}

const (
	IN_CHANNEL = "in_channel"
	EPHEMERAL  = "ephemeral"
	REPLACE_ORIGINAL = "replace_original"
	DELETE_ORIGINAL = "delete_original"
)

func NewSlackSimpleResponse(rtype string, body string) SlackSimpleResponse {
	return SlackSimpleResponse{rtype, body}
}
