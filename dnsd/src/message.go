package dnsd

import (
	"encoding/json"
)

// TopicName defines which NSQ topic requests will be published to
const TopicName string = "AddRequest"

// Publish will publish the request to NSQ
func Publish(request Request) error {

	// Serialize request to JSON
	json, error1 := json.Marshal(request)
	if error1 != nil {
		return error1
	}

	// Publish request to NSQ
	error2 := PluginProducer.Publish(TopicName, []byte(json))
	if error2 != nil {
		return error2
	}

	// Return no error
	return nil
}
