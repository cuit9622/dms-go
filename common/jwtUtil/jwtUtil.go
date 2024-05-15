package jwtUtil

import (
	"fmt"
	"strconv"

	jwt "github.com/golang-jwt/jwt/v5"
)

var secertKey = []byte("SJD(O!I@#()SKD<?X<?Z<D)P:K@_)#IO)_SI[KDL;AO)PQ@I#FKDJNFKL")

func Verify(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secertKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetUserId(tokenString string) (int64, error) {
	token, err := Verify(tokenString)
	if err != nil {
		return 0, err
	}
	subject, err := token.Claims.GetSubject()
	if err != nil {
		return 0, err
	}
	id, err := strconv.ParseInt(subject, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}
