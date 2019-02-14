package slack

import (
        "log"
        "reflect"
)

func AttachSlackoverflow(sc *Client) error {
        channels, err := sc.App.GetChannels(false)
        if err != nil {
            log.Panic(err)
        }

        for _, channel := range channels {
            log.Println(reflect.TypeOf(channel))
        }

        return nil
}
