package app

import (
	"fmt"
	"os"
	"context"

	"github.com/ondrej-smola/mesos-go-http/lib/client"
	"github.com/ondrej-smola/mesos-go-http/lib/client/leader"
	"github.com/ondrej-smola/mesos-go-http/lib/operator/master"

	"github.com/go-kit/kit/log"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/jsonpb"
	"encoding/base64"
	"strings"
)

func Run(cfg Config) error{

	w := log.NewSyncWriter(os.Stderr)
	logger := log.NewLogfmtLogger(w)

	//ctx, cancel := context.WithCancel(context.Background())
	ctx := context.Background()

	cfg.endpoints = strings.Split(cfg.url, ",")

	mesos_leader_cli := getLeader(&cfg, logger)
	//mesos_cli := master.New(mesos_leader)
	//mesos_flags, err := mesos_cli.GetFlags(ctx)
	//if err != nil {
	//	fmt.Printf("Error getting mesos master flags: %s", err)
	//	logger.Log("Error getting mesos master flags: %s", err)
	//	return err
	//}
	//printResponse(mesos_flags)

	//mesos_state, err := mesos_cli.GetLoggingLevel(ctx)
	//if err != nil {
	//	fmt.Printf("Error getting mesos master state: %s", err)
	//	logger.Log("Error getting mesos master state: %s", err)
	//	return err
	//}
	//logger.Log("State without error, printing...")
	//printResponse(mesos_state)

	fmt.Println("Creating EventStream")
	events := master.NewEventStream(mesos_leader_cli, ctx)
	fmt.Println("EventStream created")
	for event := range events {
		fmt.Println("existen eventos")
		if event.Err == nil {
			fmt.Println("Impriemiendo eventos")
			logger.Log("Event %s received, printing...", event.Event.GetType().String())
			if event.Event.GetType().String() == "TASK_ADDED" {
				fmt.Printf("Network Info: %s\n Statuses: %s\n",
					event.Event.GetTaskAdded().GetTask().GetContainer().GetNetworkInfos(),
					event.Event.GetTaskAdded().GetTask().GetStatuses())
			}
			//printResponse(event.Event)
		} else {
			logger.Log("Error in event: %s", event.Err)
		}
	}
	fmt.Println("End")
	return nil

}

func getLeader(cfg *Config, logger log.Logger) *leader.LeaderClient {

	auth := cfg.credentials.username + ":" + cfg.credentials.password
	credential := base64.StdEncoding.EncodeToString([]byte(auth))
	credential = base64.StdEncoding.EncodeToString([]byte("open:DrE6o5r67qnf38H9K5KgGhoeRHqwU6sVNBVStHqw"))

	mesos_leader := leader.New(cfg.endpoints,
		leader.WithLogger(logger),
		leader.WithClientOpts(
			client.WithRecordIOFraming(),
			client.WithRequestOpts(
				client.WithAuthorization("Basic " + credential),
				client.WithHeader("User-Agent", "registrator/0.1"),
			),
		))
	return mesos_leader
}

//func getTasks(cfg *Config) {
//	v, err := getLeader().GetTasks(cfg.ctx)
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//
//	}
//	printResponse(v)
//}

func printResponse(m proto.Message) {
	marsh := jsonpb.Marshaler{EmitDefaults: true}
	b, err := marsh.MarshalToString(m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(b)
	}
}