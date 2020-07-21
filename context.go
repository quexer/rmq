package rmq

import (
	"context"

	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
)

// setSubscribeOption returns a function to setup a context with given value
func setSubscribeOption(k, v interface{}) broker.SubscribeOption {
	return func(o *broker.SubscribeOptions) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}

// setBrokerOption returns a function to setup a context with given value
func setBrokerOption(k, v interface{}) broker.Option {
	return func(o *broker.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}

// setBrokerOption returns a function to setup a context with given value
func setServerSubscriberOption(k, v interface{}) server.SubscriberOption {
	return func(o *server.SubscriberOptions) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}

// setPublishOption returns a function to setup a context with given value
func setPublishOption(k, v interface{}) broker.PublishOption {
	return func(o *broker.PublishOptions) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}

const deliveryModeHeader = "RabbitMQ-Delivery-Mode"

// DurableMessageContext create new context with metadata for durable message sending
func DurableMessageContext(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	md[deliveryModeHeader] = "2" // 2 means durable per rabbitMQ doc
	return metadata.NewContext(ctx, md)
}
