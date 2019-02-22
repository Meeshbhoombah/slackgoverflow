package slack

import (
        "fmt"

        s "github.com/nlopes/slack"
)

func (c *Client) Ask(q string, u string) error {
        question := fmt.Sprintf(`*@%s asked:* %s`, u, q)
        c.App.PostMessage(c.ChannelId, s.MsgOptionText(question, false))
        return nil
}

func (c *Client) NotifyUser(txt string, uID string) error {
        c.App.PostEphemeral(c.ChannelId, uID, s.MsgOptionText(txt, false))
        return nil
}
