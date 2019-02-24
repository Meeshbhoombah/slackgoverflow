package slack

import (
        "fmt"

        s "github.com/nlopes/slack"
)

func (c *Client) Ask(q string, u string) error {
        user := fmt.Sprintf(`*@%s asked:*`, u)
        a := s.Attachment{
            Text: q,
        }

        c.App.PostMessage(c.ChannelId,
            s.MsgOptionText(user, false),
            s.MsgOptionAttachments(a)
        )

        return nil
}

func (c *Client) NotifyUser(txt string, userID string) error {
        c.App.PostEphemeral(c.ChannelId,
            userID,
            s.MsgOptionText(txt, false)
        )

        return nil
}
