package client

import (
	"context"
	"fmt"
	"thrift-go/gen-go/tutorial"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
)

var defaultCtx = context.Background()

func handleClient(client *tutorial.CalculatorClient) (err error) {

	start := time.Now()

	health, _ := client.Health(defaultCtx)
	fmt.Println(health)

	pong, _ := client.Ping(defaultCtx)
	fmt.Println(pong)

	sum, _ := client.Add(defaultCtx, 1, 1)
	fmt.Print("1+1=", sum, "\n")

	work := tutorial.NewWork()
	work.Op = tutorial.Operation_DIVIDE
	work.Num1 = 1
	work.Num2 = 0
	quotient, err := client.Calculate(defaultCtx, 1, work)
	if err != nil {
		switch v := err.(type) {
		case *tutorial.InvalidOperation:
			fmt.Println("Invalid operation:", v)
		default:
			fmt.Println("Error during operation:", err)
		}
	} else {
		fmt.Println("Whoa we can divide by 0 with new value:", quotient)
	}

	work.Op = tutorial.Operation_SUBTRACT
	work.Num1 = 15
	work.Num2 = 10
	diff, err := client.Calculate(defaultCtx, 1, work)
	if err != nil {
		switch v := err.(type) {
		case *tutorial.InvalidOperation:
			fmt.Println("Invalid operation:", v)
		default:
			fmt.Println("Error during operation:", err)
		}
	} else {
		fmt.Print("15-10=", diff, "\n")
	}

	log, err := client.GetStruct(defaultCtx, 1)
	if err != nil {
		fmt.Println("Unable to get struct:", err)
		return err
	} else {
		fmt.Println("Check log:", log.Value)
	}

	duration := time.Now().Sub(start).Seconds() * 1000
	fmt.Println("cost:", duration)

	return err
}

func RunClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TTransport
	var err error

	transport, err = thrift.NewTSocket(addr)

	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}

	defer transport.Close()

	if err = transport.Open(); err != nil {
		return err
	}

	protocol := protocolFactory.GetProtocol(transport)

	return handleClient(tutorial.NewCalculatorClient(thrift.NewTStandardClient(protocol, protocol)))

}
