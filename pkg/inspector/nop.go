package inspector

import "github.com/cristalhq/jwt/v5"

type JwtNopInspector struct{}

func (j *JwtNopInspector) Inspect(raw string) (ParsedToken, error) {

	r := []byte(raw)
	token, err := jwt.ParseNoVerify(r)
	if err != nil {
		//TODO: Debug logs
		return ParsedToken{}, ErrInvalidToken
	}
	out := ParsedToken{
		HeaderRaw:    string(token.HeaderPart()),
		ClaimRaw:     string(token.ClaimsPart()),
		SignatureRaw: string(token.SignaturePart()),
		Header: Header{
			string(token.Header().Algorithm),
			string(token.Header().Type),
			string(token.Header().ContentType),
		},
	}
	c := make(map[string]interface{})
	err = token.DecodeClaims(&c)
	if err != nil {
		//TODO: Debug logs
		return out, ErrInvalidClaims
	}
	out.Claims = c
	return out, nil

}
