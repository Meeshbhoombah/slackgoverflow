package slack

import (
	"fmt"
)

const (
	BaseURL = `https://slack.com/app_redirect?channel=%s&team=%s`
)

func (c *Client) GenerateDeepLink() (string, error) {
	str := fmt.Sprintf(baseURL, c.ChanID, c.TeamID)
	return str, nil
}
