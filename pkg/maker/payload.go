package maker

import (
	"errors"
	"time"
)

var (
	ErrExpiredToken   = errors.New("token has expired")
	ErrInvalidToken   = errors.New("token is invalid")
	ErrOwnerIDIsEmpty = errors.New("token owner is invalid")
)

type Payload struct {
	GameID    string    `json:"game_id"`
	OwnerID   string    `json:"owner_id"`
	IsHost    bool      `json:"is_host"`
	IssuedAt  time.Time `josn:"issuedat"`
	ExpiredAt time.Time `json:"expiredat"`
}

func NewPayload(game_id string, owner_id string, isHost bool, duration time.Duration) (*Payload, error) {
	if len(owner_id) == 0 {
		return nil, ErrOwnerIDIsEmpty
	}

	payload := &Payload{
		GameID:    game_id,
		OwnerID:   owner_id,
		IsHost:    isHost,
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
