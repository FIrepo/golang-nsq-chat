package models

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	nsq "github.com/bitly/go-nsq"
	"github.com/manhtai/golang-nsq-chat/config"
)

// NsqReader represents a NSQ channel below topic Chat
type NsqReader struct {
	channelName string
	consumer    *nsq.Consumer
	rooms       map[*Room]bool
}

// NewNsqReader create new NsqReader from a channel name
func NewNsqReader(r *Room, channelName string) error {

	cfg := nsq.NewConfig()
	cfg.Set("LookupdPollInterval", config.LookupdPollInterval*time.Second)
	cfg.Set("MaxInFlight", config.MaxInFlight)
	cfg.UserAgent = fmt.Sprintf("golang-nsq-chat go-nsq/%s", nsq.VERSION)

	nsqConsumer, err := nsq.NewConsumer(config.TopicName, channelName, cfg)

	if err != nil {
		log.Println("nsq.NewNsqReader error: ", err)
		return err
	}

	nsqReader := &NsqReader{
		channelName: channelName,
		rooms:       map[*Room]bool{r: true},
	}
	r.nsqReaders[channelName] = nsqReader

	nsqConsumer.AddHandler(nsqReader)

	nsqErr := nsqConsumer.ConnectToNSQLookupd(config.AddrNsqlookupd)
	if nsqErr != nil {
		log.Println("NSQ connection error: ", nsqErr)
		return err
	}
	nsqReader.consumer = nsqConsumer
	log.Printf("Subscribe to NSQ success to channel %s", channelName)

	return nil
}

// HandleMessage pushes messages from NSQ to Client, is used by AddHandler() function
func (nr *NsqReader) HandleMessage(msg *nsq.Message) error {
	message := Message{}
	err := json.Unmarshal(msg.Body, &message)
	if err != nil {
		log.Println("NSQ HandleMessage ERROR: invalid JSON subscribe data")
		return err
	}
	for r := range nr.rooms {
		r.forward <- &message
	}
	return nil
}

// getChannelName return sha256 hash of Hostname
func getChannelName() string {
	hostname, _ := os.Hostname()
	return fmt.Sprintf("%x", sha256.Sum256([]byte(hostname)))
}

// subscribeToNsq subscribes Room to a NSQ channel
func subscribeToNsq(r *Room) {
	nsqChannelName := getChannelName()
	_, ok := r.nsqReaders[nsqChannelName]

	if !ok {
		err := NewNsqReader(r, nsqChannelName)
		if err != nil {
			log.Printf("Failed to subscribe to channel: '%s'",
				nsqChannelName)
			return
		}
	}
}
