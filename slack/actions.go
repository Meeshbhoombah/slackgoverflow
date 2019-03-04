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

	c.Usr.PostMessage(c.ChanId,
		s.MsgOptionText(asker, false),
		s.MsgOptionAttachments(question),
	)

	return nil
}

func (c *Client) NotifyUser(txt string, userID string) error {
	c.Usr.PostEphemeral(c.ChanId,
		userID,
		s.MsgOptionText(txt, false),
	)

	return nil
}
