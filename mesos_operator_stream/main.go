package main

import (
	"context"
	"fmt"
	"os"
	"encoding/base64"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/ondrej-smola/mesos-go-http/lib/client/leader"
	"github.com/ondrej-smola/mesos-go-http/lib/operator/master"
	"github.com/ondrej-smola/mesos-go-http/lib/client"
	"github.com/go-kit/kit/log"


)

type config struct {
	endpoints []string
	printResponse func(proto.Message)
	ctx context.Context
	credentials string
}

func main() {


	w := log.NewSyncWriter(os.Stderr)
	logger := log.NewLogfmtLogger(w)

	cfg := &config{
		printResponse: jsonpbPrint,
		ctx:           context.Background(),
		endpoints:     os.Args[1:],
		credentials:   base64.StdEncoding.EncodeToString([]byte("open:DrE6o5r67qnf38H9K5KgGhoeRHqwU6sVNBVStHqw")),
	}

	getLeader := func(opts ...leader.Opt) *master.Client {
		return master.New(leader.New(
			cfg.endpoints,
			leader.WithLogger(logger),
			leader.WithClientOpts(
				client.WithRequestOpts(
					client.WithAuthorization("Basic " + cfg.credentials)))),
		)
	}

	v, err := getLeader().GetVersion(cfg.ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cfg.printResponse(v)

	events := master.NewEventStream(
		leader.New(
			cfg.endpoints,
			leader.WithClientOpts(client.WithRecordIOFraming()),
			leader.WithLogger(logger),
			leader.WithClientOpts(
				client.WithRequestOpts(
					client.WithAuthorization("Basic " + cfg.credentials))),
			),
		cfg.ctx,
	)

	fmt.Printf("Received events: %d\n", len(events))
	for ev := range events {
		fmt.Println("Pasas por loop de eventos!!")
		fmt.Println(ev)
		if ev.Err != nil {
			fmt.Println(ev.Err)
			os.Exit(1)
		} else {
			cfg.printResponse(ev.Event)
		}
	}
	fmt.Println("Fuera de loop de eventos")
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