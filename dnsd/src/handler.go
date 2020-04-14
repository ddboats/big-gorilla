package dnsd

import (
	"context"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
)

// PluginHandler will handle incoming DNS requests
type PluginHandler struct {
	Next plugin.Handler
}

// ServeDNS will handle each DNS request
func (handler PluginHandler) ServeDNS(context context.Context, responseWriter dns.ResponseWriter, message *dns.Msg) (int, error) {
	// Iterate through each DNS question the client has and add it to NSQ
	for _, question := range message.Question {
		query := CreateQuery(responseWriter, question)
		error1 := Publish(query)

		// There was an error so log it
		if error1 != nil {
			PluginLogger.Error(error1.Error())
		}
	}
	// Proceed down the plugin chain
	return plugin.NextOrFailure(PluginName, handler.Next, context, responseWriter, message)
}

// Name will return the plugin's name
func (handler PluginHandler) Name() string {
	return PluginName
}
