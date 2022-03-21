package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer c.Close()

	signalName := "stop_checking_bill"
	workflowID := "cb0fef5d-96b9-4933-9107-eb48fd525238"
	runID := "5408419d-ed18-4723-80a6-e672a5d13651"

	if err = c.SignalWorkflow(context.Background(), workflowID, runID, signalName, true); err != nil {
		log.Panicln("fail to send signal", err)
	}
}
