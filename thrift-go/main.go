package main

import (
	"flag"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
	"thrift-go/server"

	"thrift-go/client"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage

	serverFlag := flag.Bool("server", false, "Run server")

	protocol := flag.String("P", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	framed := flag.Bool("framed", false, "Use framed transport")
	buffered := flag.Bool("buffered", false, "Use buffered transport")
	addr := flag.String("addr", "localhost:9090", "Address to listen to")

	flag.Parse()

	var protocalFactory thrift.TProtocolFactory
	switch *protocol {
	case "compact":
		protocalFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocalFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocalFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocalFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol, "\n")
		Usage()
		os.Exit(1)
	}

	var transportFactory thrift.TTransportFactory
	if *buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	if *framed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}

	if *serverFlag {
		if err := server.RunServer(transportFactory, protocalFactory, *addr); err != nil {
			fmt.Println("error running server:", err)
		}
	} else {
		if err := client.RunClient(transportFactory, protocalFactory, *addr); err != nil {
			fmt.Println("error running client:", err)
		}
	}
}
