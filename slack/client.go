package slack

import (
	"bytes"
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
	*s.Client

	// Match w/ VerificationToken on incoming request to verify
	Ver string

	// Token of user who installed Slackoverflow
	Usr string

	// Workspace Team ID
	TeamId string
	// Workspace Team Name
	TeamName string
	// `#slackoverflow` Channel ID
	ChanId string
}

func Init(cfg *config.Variables, db *gorm.DB, accCode *string) (*Client, error) {
	// Use access code returned from Slack Authorization to get credentials
	b := url.Values{
		"client_id":     cfg.SlackClientId,
		"client_secret": cfg.SlackClientSecret,
		"code":          accCode,
		"redirect_uri":  redirectUri,
	}

	body := bytes.NewBufferString(b.Encode)

	r, err := http.Post(baseURL, "application/x-www-form-urlencoded", body)
	if err != nil {
		log.Error(err)
	}
	defer r.Body.Close()

	rsp := new(s.OAuthResponse)

	// `nlopes/slack` provides binding for OAuth response
	err := json.NewDecoder(r.Body).Decode(rsp)
	if err != nil {
		log.Error(err)
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
