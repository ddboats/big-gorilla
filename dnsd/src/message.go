package dnsd

import (
	"time"
	"encoding/json"

	"github.com/nsqio/go-nsq"
)

// TopicName defines which NSQ topic queries will be published to
const TopicName string = "AddQuery"

// Publish will publish the query to NSQ
func Publish(query Query) error {

	// Serialize query to JSON
	json, error1 := json.Marshal(query)
	if error1 != nil {
		return error1
	}

	// Handle the producer's transaction when it is complete
	doneChannel := make(chan *nsq.ProducerTransaction)
	go func() {
		transaction := <-doneChannel
		if transaction.Error != nil {
			PluginLogger.Error(transaction.Error.Error())
		}
	}()

	// Publish request to NSQ
	error2 := PluginProducer.DeferredPublishAsync(TopicName, time.Second, []byte(json), doneChannel)
	if error2 != nil {
		return error2
	}

	// Return no error
	return nil
}
