package storage

import (
	"errors"
	"log"
	"telegram-discord-pipe-bot/interfaces"
	"telegram-discord-pipe-bot/models"
)

var (
	errNoChannelRegistered = errors.New("no channel with given id is registered")
	errAlreadyRegistered   = errors.New("the given channel is already registered")
)

type Memory struct {
	registeredChannels []models.CommunicationChannel
}

func NewInmemory() *Memory {
	return &Memory{}
}

func (m *Memory) Register(commChannelID, commChannelName string) error {
	m.registeredChannels = append(m.registeredChannels, models.CommunicationChannel{
		ChannelID:   commChannelID,
		ChannelName: commChannelName,
	})
	log.Printf("in memory channel %s is registered", commChannelName)

	return nil
}

func (m *Memory) Unregister(commChannelID string) error {
	for i, c := range m.registeredChannels {
		if c.ChannelID == commChannelID {
			m.registeredChannels = append(m.registeredChannels[:i], m.registeredChannels[i+1:]...)
			log.Printf("in memory channel %s is unregistered", c.ChannelName)

			return nil
		}
	}

	return errNoChannelRegistered
}

func (m *Memory) PrePopulateCommunicationList(commList []models.CommunicationChannel) error {
	m.registeredChannels = commList
	return nil
}

func (m *Memory) PopulateCommunicationList() error {
	m.registeredChannels = make([]models.CommunicationChannel, 0)
	log.Println("in memory channel list is initialized")
	return nil
}

func (m *Memory) GetCommunicationList() ([]models.CommunicationChannel, error) {
	return m.registeredChannels, nil
}

func (m *Memory) IsRegistered(commChannelID string) (bool, error) {
	for _, c := range m.registeredChannels {
		if c.ChannelID == commChannelID {
			return true, nil
		}
	}

	return false, nil
}

var _ interfaces.CommunicationStorager = (*Memory)(nil)
