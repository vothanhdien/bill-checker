package main

import (
	"log"

	ia "github.com/vothanhdien/bill-checker/internal/activity"
	"github.com/vothanhdien/bill-checker/internal/constant"
	iw "github.com/vothanhdien/bill-checker/internal/workflow"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, constant.CheckBillTaskQueue, worker.Options{})
	w.RegisterWorkflow(iw.CheckBillWorkFlow)

	// register activity
	act := &ia.BillChecker{
		Checker: &ia.CheckerStub{M: make(map[string]int)},
	}

	w.RegisterActivityWithOptions(act.QueryBill, activity.RegisterOptions{Name: "QueryBill"})

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Panicln("unable to start worker", err)
	}
}
