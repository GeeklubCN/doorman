package state

import (
	"encoding/base64"
)

type State interface {
	Encode(rawUrl string) string
	Decode(escapedUrl string) (string, error)
}

type SimpleState struct{}

func (s SimpleState) Encode(rawUrl string) string {
	return base64.URLEncoding.EncodeToString([]byte(rawUrl))
}

func (s SimpleState) Decode(escapedUrl string) (string, error) {
	res, err := base64.URLEncoding.DecodeString(escapedUrl)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
