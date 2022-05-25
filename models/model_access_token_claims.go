package models

import "github.com/golang-jwt/jwt"

type AccessTokenClaims struct {
	Iss   string      `json:"iss" yaml:"iss" bson:"iss" mapstructure:"Iss"`
	Sub   string      `json:"sub" yaml:"sub" bson:"sub" mapstructure:"Sub"`
	Aud   interface{} `json:"aud" yaml:"aud" bson:"aud" mapstructure:"Aud"`
	Scope string      `json:"scope" yaml:"scope" bson:"scope" mapstructure:"Scope"`
	Exp   int32       `json:"exp" yaml:"exp" bson:"exp" mapstructure:"Exp"`
	jwt.StandardClaims
}
