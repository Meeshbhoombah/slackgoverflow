package slack

import (
	"fmt"

	s "github.com/nlopes/slack"
)

func (c *Client) Ask(q string, usr string) error {
	fmt.Println("Asking question @ ", c.ChanID)
	fmt.Println(usr)

	asker := fmt.Sprintf(`<@%s> *asked:*`, usr)
	question := s.Attachment{
		Text: q,
	}

	rsp, _, err := c.Usr.PostMessage(c.ChanID,
		s.MsgOptionText(asker, false),
		s.MsgOptionAttachments(question),
	)

	fmt.Println(rsp)
	fmt.Println(err)

	return nil
}

func (c *Client) NotifyUser(txt string, userID string) error {
	c.Usr.PostEphemeral(c.ChanID,
		userID,
		s.MsgOptionText(txt, false),
	)

	return nil
}
