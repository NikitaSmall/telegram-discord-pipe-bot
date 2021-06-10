package interfaces

import "telegram-discord-pipe-bot/models"

type CommunicationStorager interface {
	Register(commChannelID, commChannelName string) error
	Unregister(commChannelID string) error

	PopulateCommunicationList() error

	GetCommunicationList() ([]models.CommunicationChannel, error)
	IsRegistered(commChannelID string) (bool, error)
}
