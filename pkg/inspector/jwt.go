// Package inspector provides functionality for introspecting json-web-tokens 
package inspector

import "errors"

var (
	ErrInvalidClaims = errors.New("claims are invalid")
	ErrInvalidToken  = errors.New("invalid token")
)

type Inspector interface {
	Inspect(token string) (ParsedToken, error)
}

// ParsedToken is the ouput of this package 
type ParsedToken struct {
	HeaderRaw    string
	ClaimRaw     string
	SignatureRaw string

	Header Header
	Claims map[string]interface{}
}

// Headers represents the token header information 
type Header struct {
	Algorithm   string
	Type        string
	ContentType string
}
