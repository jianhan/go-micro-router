package middleware

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"os"
	"github.com/y0ssar1an/q"
	"errors"
	"github.com/jianhan/go-micro-router/handler"
)

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options {





	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Verify 'aud' claim
			aud := os.Getenv("AUTH0_AUDIENCE")
			q.Q(token.Claims.(jwt.MapClaims))
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !checkAud {
				return token, errors.New("Invalid audience.")
			}
			// Verify 'iss' claim
			iss := "https://" + os.Getenv("AUTH0_DOMAIN") + "/"
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return token, errors.New("Invalid issuer.")
			}

			cert, err := handler.GetPemCert(token)
			if err != nil {
				panic(err.Error())
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})