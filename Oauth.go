package freckleapi

import (
	"golang.org/x/oauth2"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://secure.letsfreckle.com/oauth/2/authorize",
	TokenURL: "https://secure.letsfreckle.com/oauth/2/access_token",
}
