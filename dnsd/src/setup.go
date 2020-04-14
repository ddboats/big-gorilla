package dnsd

import (
	"errors"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"

	"github.com/caddyserver/caddy"

	"github.com/nsqio/go-nsq"
)

// PluginName is the name of this plugin
const PluginName string = "big-gorilla"

// PluginProducer is the NSQ producer
var PluginProducer *nsq.Producer = nil

// init will begin plugin initialization
func init() {
	// Register the plugin with CoreDNS
	plugin.Register(PluginName, setup)
}

// setup will register the plugin with CoreDNS and connect to NSQ
func setup(controller *caddy.Controller) error {
	// Proceed to first token
	controller.Next()

	// Determine if there's any unnecessary arguments
	if controller.NextArg() {
		// There was so return an error
		return plugin.Error(PluginName, controller.ArgErr())
	}

	// When the CoreDNS server starts initialize the NSQ connection
	controller.OnStartup(func() error {
		producer, error1 := initNSQ()
		PluginProducer = producer
		return error1
	})

	// When the CoreDNS server shuts down close the NSQ connection
	controller.OnShutdown(func() error {
		if PluginProducer == nil {
			return errors.New("nsq producer nil")
		}
		PluginProducer.Stop()
		return nil
	})

	// Tell CoreDNS to use this plugin
	dnsserver.GetConfig(controller).AddPlugin(func(handler plugin.Handler) plugin.Handler {
		return PluginHandler{}
	})

	// Return no errors
	return nil
}

// initNSQ will attempt to create an NSQ producer
func initNSQ() (*nsq.Producer, error) {
	config := nsq.NewConfig()
	return nsq.NewProducer("nsqd:4150", config)
}
