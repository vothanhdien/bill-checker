package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vothanhdien/bill-checker/internal/constant"
	iw "github.com/vothanhdien/bill-checker/internal/workflow"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "foo",
		TaskQueue: constant.CheckBillTaskQueue,
	}
	we, err := c.ExecuteWorkflow(context.Background(), options, iw.CheckBillWorkFlow, &iw.CheckBillWFInput{CusCode: ""})
	if err != nil {
		log.Panicln("unable to start workflow executions", err)
	}

	var out iw.CheckBillWFOutput
	err = we.Get(context.Background(), &out)
	if err != nil {
		log.Panicln("unable to get workflow result", err)
	}
	printResults(out, we.GetID(), we.GetRunID())
}

func printResults(out iw.CheckBillWFOutput, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", out)
}
