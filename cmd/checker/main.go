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
		TaskQueue: constant.CheckBillTaskQueue,
	}

	in := &iw.CheckBillWFInput{
		CusCode: "FOO_1000",
	}
	we, err := c.ExecuteWorkflow(context.Background(), options, iw.CheckBillWorkFlow, in)
	if err != nil {
		log.Panicln("unable to start workflow executions", err)
	}

	var out iw.CheckBillWFOutput
	err = we.Get(context.Background(), &out)
	if err != nil {
		log.Panicln("unable to get workflow result", err)
	}
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())
	fmt.Printf("\nCustomerCode:%s\n", in.CusCode)
	for i, v := range out.Attempts {
		fmt.Printf("%v - %v - %v\n", i, v.T, v.B)
	}
}
