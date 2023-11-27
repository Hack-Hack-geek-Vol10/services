package maker

import (
	"errors"
	"time"

	token "github.com/Hack-Hack-geek-Vol10/services/pkg/grpc/token-service/v1"
)

var (
	ErrExpiredToken     = errors.New("token has expired")
	ErrInvalidToken     = errors.New("token is invalid")
	ErrProjectIDIsEmpty = errors.New("project id is empty")
)

type Payload struct {
	ProjectID string     `json:"game_id"`
	Authority token.Auth `json:"authority"`
	IssuedAt  time.Time  `josn:"issuedat"`
	ExpiredAt time.Time  `json:"expiredat"`
}

func NewPayload(projectID string, authority token.Auth, duration time.Duration) (*Payload, error) {
	if len(projectID) == 0 {
		return nil, ErrProjectIDIsEmpty
	}

	payload := &Payload{
		ProjectID: projectID,
		Authority: authority,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
