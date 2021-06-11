package models

type CommunicationChannel struct {
	ChannelID   string `json:"channelId" firestore:"channelId"`
	ChannelName string `json:"channelName" firestore:"channelName"`
}
