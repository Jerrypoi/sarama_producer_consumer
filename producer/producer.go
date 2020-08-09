package main

import (
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
)

var (
	dst string
)
func parseArgs() {
	flag.StringVar(&dst,"destination", "localhost:9092", "destination of the server")

}

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	fmt.Printf("Going to init connection")

	client, err := sarama.NewSyncProducer([]string{dst}, config)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer client.Close()
	fmt.Printf("Init client Succ")
	msg := &sarama.ProducerMessage{}
	msg.Topic = "TutorialTopic"
	msg.Value = sarama.StringEncoder("Hello from mbp")
	pid, offset ,err := client.SendMessage(msg)

	if err != nil {
		fmt.Printf("Error!!")
		os.Exit(-1)
	}

	fmt.Printf("pid: %v, offset: %v\n", pid, offset)

}