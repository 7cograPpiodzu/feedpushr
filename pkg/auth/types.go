package auth

import "net/http"

// Authenticator is a generic interface to validate an HTTP request
type Authenticator interface {
	Validate(req *http.Request, res http.ResponseWriter) bool
	Issuer() string
}
