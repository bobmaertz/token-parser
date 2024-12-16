package inspector

import "errors"

var (
	ErrInvalidClaims = errors.New("claims are invalid")
	ErrInvalidToken  = errors.New("invalid token")
)

type Inspector interface {
	Inspect(token string) (ParsedToken, error)
}

// JWT is the ouput of this package. This is more complex than i like
type ParsedToken struct {
	HeaderRaw    string
	ClaimRaw     string
	SignatureRaw string

	Header Header
	Claims map[string]interface{}
}

type Header struct {
	Algorithm   string
	Type        string
	ContentType string
}
