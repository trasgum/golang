package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/ondrej-smola/mesos-go-http/lib/client/leader"
	"github.com/ondrej-smola/mesos-go-http/lib/operator/master"
	"github.com/ondrej-smola/mesos-go-http/lib/client"
	"encoding/base64"
)

type config struct {
	endpoints []string
	printResponse func(proto.Message)
	ctx context.Context
	credentials []string
}

func main() {
	cfg := &config{
		printResponse: jsonpbPrint,
		ctx:           context.Background(),
		endpoints:     os.Args[1:],
		credentials:   base64.StdEncoding.EncodeToString([]byte("open:DrE6o5r67qnf38H9K5KgGhoeRHqwU6sVNBVStHqw")),
	}

	leader.New(cfg.endpoints,
		leader.WithClientOpts(client.WithRecordIOFraming()),
	)

	leader.WithClientOpts(client.WithAuthorization(cfg.credentials))
	leader.Opt(client.WithAuthorization(cfg.credentials))

	events := master.NewEventStream(
		leader.New(cfg.endpoints,
			leader.WithClientOpts(client.WithRecordIOFraming()),
			),
		cfg.ctx,
	)

	for ev := range events {
				fmt.Println(ev)
				if ev.Err != nil {
					fmt.Println(ev.Err)
					os.Exit(1)
				} else {
					cfg.printResponse(ev.Event)
				}

			}

}

func jsonpbPrint(m proto.Message) {
		marsh := jsonpb.Marshaler{EmitDefaults: true}
	b, err := marsh.MarshalToString(m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(b)
	}
}