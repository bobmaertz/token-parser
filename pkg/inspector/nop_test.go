package inspector

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

func TestJwtNopInspector_Inspect(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		j       *JwtNopInspector
		args    args
		want    ParsedToken
		wantErr error
	}{
		{
			name: "Successfully extract jwt data",
			j:    &JwtNopInspector{},
			args: args{
				//Dummy token pulled from jwt.io
				raw: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			},
			want: ParsedToken{
				HeaderRaw:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
				ClaimRaw:     "eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ",
				SignatureRaw: "SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
				Header: Header{
					Algorithm:   "HS256",
					Type:        "JWT",
					ContentType: "",
				},
				Claims: map[string]interface{}{
					"iat": func() float64 {
						str := "1.516239022e+09"
						num, err := strconv.ParseFloat(str, 64)
						if err != nil {
							t.Errorf("Test setup issue: %v", err)
						}
						return num
					}(),
					"name": "John Doe",
					"sub":  "1234567890",
				},
			},
			wantErr: nil,
		},
		{
			name: "Token is invalid",
			j:    &JwtNopInspector{},
			args: args{
				//Dummy token pulled from jwt.io
				raw: "FAKEeyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			},
			want:    ParsedToken{},
			wantErr: ErrInvalidToken,
		},
		{
			name: "Claims are invalid",
			j:    &JwtNopInspector{},
			args: args{
				//Dummy token pulled from jwt.io
				// Error introduced with XXXFAKEXXX
				raw: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIXXXFAKEXXXDkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			},
			want: ParsedToken{
				HeaderRaw:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
				ClaimRaw:     "eyJzdWIiOiIXXXFAKEXXXDkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ",
				SignatureRaw: "SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
				Header: Header{
					Algorithm:   "HS256",
					Type:        "JWT",
					ContentType: "",
				},
				Claims: nil,
			},
			wantErr: ErrInvalidClaims,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JwtNopInspector{}
			got, err := j.Inspect(tt.args.raw)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("JwtNopInspector.Inspect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JwtNopInspector.Inspect() = %v, want %v", got, tt.want)
			}
		})
	}
}
