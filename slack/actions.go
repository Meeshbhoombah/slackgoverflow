package slack

import (
	"fmt"

	s "github.com/nlopes/slack"
)

func (c *Client) Ask(q string, usr string) error {
	asker := fmt.Sprintf(`<@%s> *asked:*`, usr)
	question := s.Attachment{
		Text: q,
	}

	c.App.PostMessage(c.ChannelId,
		s.MsgOptionText(asker, false),
		s.MsgOptionAttachments(question),
	)

	return nil
}

func (c *Client) NotifyUser(txt string, userID string) error {
	c.App.PostEphemeral(c.ChannelId,
		userID,
		s.MsgOptionText(txt, false),
	)

	return nil
}
