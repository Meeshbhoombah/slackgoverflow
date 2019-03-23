package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/jinzhu/gorm"
	s "github.com/nlopes/slack"

	"github.com/archproj/slackoverflow/config"
	"github.com/archproj/slackoverflow/models"
)

const (
	baseURL     = `slack.com/api/oauth.access`
	redirectUri = `slackoverflow.app/`
)

type Client struct {
	// Match w/ VerificationToken on incoming request to verify
	Ver string

	// Token of user who installed Slackoverflow
	Usr *s.Client

	// Workspace Team ID
	TeamID string
	// Workspace Team Name
	TeamName string
	// Name of Channel in which Slackoverflow exists
	ChanName string
	// Channel ID
	ChanID string
}

func Init(cfg *config.Variables, db *gorm.DB, accCode *string) (*Client, error) {
	v := url.Values{}

	v.Set("client_id", cfg.SlackClientID)
	v.Set("client_secret", cfg.SlackClientSecret)
	v.Set("code", *accCode)
	v.Set("redirect_uri", redirectUri)

	b := v.Encode()
	body := bytes.NewBufferString(b)

	// Use access code returned from Slack Authorization to get credentials
	r, err := http.Post(baseURL, "application/x-www-form-urlencoded", body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	rsp := new(s.OAuthResponse)

	// `nlopes/slack` provides binding for OAuth response
	err = json.NewDecoder(r.Body).Decode(rsp)
	if err != nil {
		return nil, err
	}

	// Persist team now, for end-user installation suspense
	// TODO: remove suspense w/ concurrency
	w := models.Workspace{
		TeamName:  rsp.TeamName,
		TeamID:    rsp.TeamID,
		UserToken: rsp.AccessToken,
		ChanName:  rsp.IncomingWebhook.Channel,
		ChanID:    rsp.IncomingWebhook.ChannelID,
	}

	db.Create(&w)

	sc, err := NewClient(cfg, &w)
	if err != nil {
		return nil, err
	}

	return sc, nil
}

func NewClient(cfg *config.Variables, w *models.Workspace) (*Client, error) {
	sc := Client{
		Ver:      cfg.SlackVerToken,
		Usr:      s.New(w.UserToken),
		TeamName: w.TeamName,
		TeamID:   w.TeamID,
		ChanName: w.ChanName,
		ChanID:   w.ChanID,
	}

	return &sc, nil
}
