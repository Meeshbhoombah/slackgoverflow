package slack

import (
	"testing"
)

var cases = []struct {
	TeamID string
	ChanID string
	Want   string
}{
	{"a", "b", fmt.Sprintf},
}

func TestGenerateDeepLink(t *testing.T) {
	for _, tc := range cases {
		log.Println("test")
	}
}
