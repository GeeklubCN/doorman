package state

import "net/url"

type State interface {
	Encode(rawUrl string) string
	Decode(escapedUrl string) (string, error)
}

type SimpleState struct{}

func (s SimpleState) Encode(rawUrl string) string {
	return url.PathEscape(rawUrl)
}

func (s SimpleState) Decode(escapedUrl string) (string, error) {
	return url.PathUnescape(escapedUrl)
}
