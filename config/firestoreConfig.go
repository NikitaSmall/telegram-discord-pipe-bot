package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type FirestoreConfig struct {
	ProjectID   string `env:"PROJECT_ID" env-required:"true"`
	Credentials string `env:"GOOGLE_APPLICATION_CREDENTIALS" env-required:"true"`
}

func (dc FirestoreConfig) GetProjectID() string {
	return dc.ProjectID
}

var firestoreConfig FirestoreConfig

func GetFirestoreConfig() FirestoreConfig {
	return firestoreConfig
}

func initFirestore() {
	log.Println("parsing configuration for firestore")
	if err := cleanenv.ReadEnv(&firestoreConfig); err != nil {
		log.Fatalf("error parsing configuration for firestore: %e", err)
	}
}
