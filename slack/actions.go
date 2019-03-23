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

	c.Usr.PostMessage(c.ChanID,
		s.MsgOptionText(asker, false),
		s.MsgOptionAttachments(question),
	)

	return nil
}

func (c *Client) NotifyUser(txt string, userID string) error {
	c.Usr.PostEphemeral(c.ChanID,
		userID,
		s.MsgOptionText(txt, false),
	)

	return nil
}
