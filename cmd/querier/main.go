package main

import (
	"context"
	"fmt"
	"log"

	iw "github.com/vothanhdien/bill-checker/internal/workflow"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer c.Close()

	queryType := "QueriedHistory"
	workflowID := "cb0fef5d-96b9-4933-9107-eb48fd525238"
	runID := "5408419d-ed18-4723-80a6-e672a5d13651"

	response, err := c.QueryWorkflow(context.Background(), workflowID, runID, queryType)
	if err != nil {
		fmt.Println("Error querying workflow: " + err.Error())
		return
	}

	var attempts []iw.CheckBillWFAttempt
	if err = response.Get(&attempts); err != nil {
		log.Panicln("query error", err)
	}
	for i, v := range attempts {
		fmt.Printf("%v - %v - %v\n", i, v.T, v.B)
	}
}
