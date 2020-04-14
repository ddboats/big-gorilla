package dnsd

import (
	"context"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"

	clog "github.com/coredns/coredns/plugin/pkg/log"
)

// log is used to log messages from the plugin to CoreDNS
var log = clog.NewWithPlugin(PluginName)

// PluginHandler will handle incoming DNS requests
type PluginHandler struct {
	Next plugin.Handler
}

// ServeDNS will handle each DNS request
func (handler PluginHandler) ServeDNS(context context.Context, responseWriter dns.ResponseWriter, message *dns.Msg) (int, error) {
	// Iterate through each DNS question the client has and add it to NSQ
	for _, question := range message.Question {
		request := CreateRequest(responseWriter, question)
		error1 := Publish(request)

		// There was an error so log it
		if error1 != nil {
			log.Error(error1.Error())
		}
	}
	// Proceed down the plugin chain
	return plugin.NextOrFailure(PluginName, handler.Next, context, responseWriter, message)
}

// Name will return the plugin's name
func (handler PluginHandler) Name() string {
	return PluginName
}
