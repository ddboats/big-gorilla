package dnsd

import (
	"encoding/json"
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

	// Publish query to NSQ
	error2 := PluginProducer.Publish(TopicName, []byte(json))
	if error2 != nil {
		return error2
	}

	// Return no error
	return nil
}
