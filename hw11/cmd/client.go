package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dark705/otus/hw11/internal/config"
	"github.com/dark705/otus/hw11/internal/logger"
	"github.com/dark705/otus/hw11/pkg/calendar/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var cFile string
	flag.StringVar(&cFile, "config", "config/config.yaml", "Config file")
	flag.Parse()
	if cFile == "" {
		_, _ = fmt.Fprint(os.Stderr, "Not set config file")
		os.Exit(2)
	}

	conf, err := config.ReadFromFile(cFile)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}

	log := logger.GetLogger(conf)
	defer logger.CloseLogFile()
	_ = log

	opts := []grpc.DialOption{grpc.WithInsecure()}

	conn, err := grpc.Dial(conf.GrpcListen, opts...)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}
	client := protobuf.NewCalendarClient(conn)
	ctx := context.TODO()

	//Event0
	_, err = client.AddEvent(ctx, &protobuf.Event{StartTime: 1000, EndTime: 2000, Title: "title1", Description: "description1"})
	fmt.Println(err)

	//Event1
	_, err = client.AddEvent(ctx, &protobuf.Event{StartTime: 2000, EndTime: 3000, Title: "title2", Description: "description2"})
	fmt.Println(err)

	//Event2
	_, err = client.AddEvent(ctx, &protobuf.Event{StartTime: 1000, EndTime: 4000, Title: "title3", Description: "description3"})
	fmt.Println(err)

	grpcEvent, err := client.GetEvent(ctx, &protobuf.Id{Id: 0})
	fmt.Println(grpcEvent, err)

	grpcEvents, err := client.GetAllEvents(ctx, &empty.Empty{})
	fmt.Println(grpcEvents, err)

	_, _ = client.DelEvent(ctx, &protobuf.Id{Id: 0})

	grpcEvents2, err := client.GetAllEvents(ctx, &empty.Empty{})
	fmt.Println(grpcEvents2, err)

}
