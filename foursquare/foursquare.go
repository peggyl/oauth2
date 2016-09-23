package foursquare

import "golang.org/x/oauth2"

// Endpoint is Foursquare's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://foursquare.com/oauth2/authenticate",
	TokenURL: "https://foursquare.com/oauth2/access_token",
}
