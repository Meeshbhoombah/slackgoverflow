
package config

import (
    "testing"
)

func TestLoadFrom(t *testing.T) {
    var c Config

    err := LoadFrom("/test.env"); if err != nil {
        t.Error(err)
    }

    if c.SecretKey != "one" {
        t.Error("Expected SECRET_KEY value to be 'one', got", c.SecretKey)
    }

    if c.Dbhost != "five" {
        t.Error("Expected DBHOST value to be 'five', got", c.Dbhost)
    }

    if c.SlackBotToken != "eleven" {
        t.Error("Expect SLACK_BOT_TOKEN value to be 'eleven', got", c.SlackBotToken)
    }
}

func TestLoad(t *testing.T) {
    load :=
}


