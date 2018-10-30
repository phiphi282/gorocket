package rest

import (
	"net/http"

	"github.com/phiphi282/gorocket/api"
)

type imsResponse struct {
	Success  bool          `json:"success"`
	Channels []api.Channel `json:"ims"`
}

// Returns all direct messages that the user has joined.
//
// https://rocket.chat/docs/developer-guides/rest-api/im/list
func (c *Client) GetJoinedIMs() ([]api.Channel, error) {
	request, _ := http.NewRequest("GET", c.getUrl()+"/api/v1/im.list", nil)
	response := new(imsResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Channels, nil
}
