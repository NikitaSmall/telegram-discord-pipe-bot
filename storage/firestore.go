package storage

import (
	"context"
	"log"
	"sync"
	"telegram-discord-pipe-bot/config"
	"telegram-discord-pipe-bot/interfaces"
	"telegram-discord-pipe-bot/models"

	"cloud.google.com/go/firestore"
)

type Firestore struct {
	inMemory *Memory
	client   *firestore.Client

	collectionName string
}

var (
	fClientInitializer sync.Once
	firestoreClient    *firestore.Client
)

func newFirestoreClient(projectID string) (*firestore.Client, error) {
	client, err := firestore.NewClient(context.Background(), projectID)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewFirestore(collectionName string) *Firestore {
	fClientInitializer.Do(func() {
		client, err := newFirestoreClient(config.GetFirestoreConfig().GetProjectID())
		if err != nil {
			log.Panicf("unable to initialize firestore client with err: %e", err)
		}

		log.Println("firestore client connection is established")
		firestoreClient = client
	})

	return &Firestore{
		inMemory:       NewInmemory(),
		client:         firestoreClient,
		collectionName: collectionName,
	}
}

func (f *Firestore) Register(commChannelID, commChannelName string) error {
	if isReg, _ := f.IsRegistered(commChannelID); isReg {
		return errAlreadyRegistered
	}

	_, err := f.client.Collection(f.collectionName).Doc(commChannelID).Set(context.Background(), models.CommunicationChannel{
		ChannelID:   commChannelID,
		ChannelName: commChannelName,
	})

	if err != nil {
		return err
	}

	log.Printf("in firestore channel %s is registered", commChannelName)
	return f.inMemory.Register(commChannelID, commChannelName)
}

func (f *Firestore) Unregister(commChannelID string) error {
	if isReg, _ := f.IsRegistered(commChannelID); !isReg {
		return errNoChannelRegistered
	}

	_, err := f.client.Collection(f.collectionName).Doc(commChannelID).Delete(context.Background())
	if err != nil {
		return err
	}

	log.Printf("in firestore channel %s is unregistered", commChannelID)
	return f.inMemory.Unregister(commChannelID)
}

func (f *Firestore) PopulateCommunicationList() error {
	commList := make([]models.CommunicationChannel, 0)
	docs, err := f.client.Collection(f.collectionName).Documents(context.Background()).GetAll()
	if err != nil {
		log.Printf("cannot get stored objects in firestore because of err: %e", err)
		return err
	}

	for _, docRef := range docs {
		var commChan models.CommunicationChannel
		if err := docRef.DataTo(&commChan); err != nil {
			log.Printf("cannot parse stored object in firestore because of err: %e", err)
			continue
		}

		commList = append(commList, commChan)
	}

	log.Printf("in firestore channels for %s group are collected. Amount %d", f.collectionName, len(commList))
	f.inMemory.PrePopulateCommunicationList(commList)
	return nil
}

func (f *Firestore) GetCommunicationList() ([]models.CommunicationChannel, error) {
	return f.inMemory.GetCommunicationList()
}

func (f *Firestore) IsRegistered(commChannelID string) (bool, error) {
	return f.inMemory.IsRegistered(commChannelID)
}

var _ interfaces.CommunicationStorager = (*Firestore)(nil)
