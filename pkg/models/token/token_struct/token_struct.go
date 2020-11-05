package token_struct

import "time"

const (
	TokenTypeOAuth  = "oauth"
	TokenTypeClient = "client"
)

const (
	ClientTypeWeb      = "web"
	ClientTypeMobile   = "mobile"
	ClientTypePersonal = "personal"
	ClientTypeOther    = "other"
)

type Token struct {
	ID         string    `graphql:"id"`
	Token      string    `graphql:"token"`
	Scopes     []string  `graphql:"scopes"`
	Valid      bool      `graphql:"valid"`
	CreatedOn  time.Time `graphql:"created_on"`
	TokenType  string    `graphql:"type"`
	UserID     string    `graphql:"user_id"`
	ClientType string    `graphql:"client_type"`
}
