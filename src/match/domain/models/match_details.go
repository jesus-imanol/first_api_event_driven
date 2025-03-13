package models

import "apiInvitation/src/match/domain/entities"

type MatchWithDetails struct {
	entities.Match
	SenderName      string `json:"sender_name"`
	ReceiverName    string `json:"receiver_name"`
	SenderPicture   string `json:"sender_picture"`
	ReceiverPicture string `json:"receiver_picture"`
}