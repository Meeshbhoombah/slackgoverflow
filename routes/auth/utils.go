package auth

import (
	"net/url"

	"github.com/archproj/slackoverflow/config"
)

const (
	// allows easy assembly of URL, permanent oauth endpoint
	BaseURL = `http://www.slack.com/oauth/authorize?`
)

var (
	// allows requisiton of all necessary scopes at time of authentication
	Scopes = [...]string{
		"incoming-webhook",
	}
)

func GenerateURL(cfg *config.Variables) (string, error) {
	// TODO: error if params are not valid
	params := url.Values{}

	params.Set("client_id", cfg.SlackClientId)
	for _, s := range Scopes {
		params.Set("scope", s)
	}
	params.Set("redirect_uri", cfg.SlackRedirectURI)

	return BaseURL + params.Encode(), nil
}
