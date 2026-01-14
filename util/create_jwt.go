package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
type Payload struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsDone    bool   `json:"is_done"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	byteArrayHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerB56 := base64UrlEncode(byteArrayHeader)

	byteArrayData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	Payloadbase64 := base64UrlEncode(byteArrayData)
	message := headerB56 + "." + Payloadbase64

	byteArraySecret := []byte(secret)
	byteArrayMessage := []byte(message)

	h := hmac.New(sha256.New, byteArraySecret)
	h.Write(byteArrayMessage)
	signature := h.Sum(nil)
	signatureB256 := base64UrlEncode(signature)
	jwt := headerB56 + "." + Payloadbase64 + "." + signatureB256
	return jwt, nil

}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)

}
