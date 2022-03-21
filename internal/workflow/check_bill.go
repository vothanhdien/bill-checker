package workflow

import (
	"time"

	"github.com/vothanhdien/bill-checker/internal/activity"
	"go.temporal.io/sdk/workflow"
)

type CheckBillWFInput struct {
	CusCode   string
	TimeRange []int
	DayRange  []int
}

type CheckBillWFAttempt struct {
	T time.Time
	B bool
}

type CheckBillWFOutput struct {
	Attempts []CheckBillWFAttempt
}

func CheckBillWorkFlow(ctx workflow.Context, in *CheckBillWFInput) (string, error) {
	customerCode := in.CusCode
	isStop := false

	var attempts []CheckBillWFAttempt

	logger := workflow.GetLogger(ctx)

	// Define query handlers
	if err := workflow.SetQueryHandler(ctx, "QueriedHistory", func() ([]CheckBillWFAttempt, error) {
		return attempts, nil
	}); err != nil {
		logger.Error("query queried history fail", err)
		return "Error", err
	}

	if err := workflow.SetQueryHandler(ctx, "CustomerCode", func() (string, error) {
		return customerCode, nil
	}); err != nil {
		logger.Error("query customerCode fail", err)
		return "Error", err
	}
	// end defining query handlers

	// Define signal channels
	// stop workflow
	selector := workflow.NewSelector(ctx)
	stopChan := workflow.GetSignalChannel(ctx, "stop_checking_bill")
	selector.AddReceive(stopChan, func(ch workflow.ReceiveChannel, _ bool) {
		logger.Info("receive signal stop")
		var b bool
		ch.Receive(ctx, &b)
		isStop = b
	})
	// end defining signal channels

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)
	logger.Info("workflow started for: " + customerCode)

	// begin
	for {
		// try to query bill
		// Charge customer for the billing period
		var result activity.QueryResult
		err := workflow.ExecuteActivity(ctx, "QueryBill", activity.QueryArgs{CusCode: in.CusCode}).Get(ctx, &result)
		if err != nil {
			logger.Error("Dont know what to do")
			return "Error", nil
		}
		attempts = append(attempts, CheckBillWFAttempt{
			T: time.Now(),
			B: result.IsHasBill,
		})

		//
		// wait
		workflow.AwaitWithTimeout(ctx, time.Minute, func() bool {
			workflow.Sleep(ctx, time.Minute)
			return isValidTime(in.TimeRange, in.DayRange)
		})

		if checkStopCondition(in.TimeRange, in.DayRange) {
			logger.Info("Over time range")
			break
		}

		if isStop {
			logger.Info("workflow stopped")
			break
		}
	}

	return "Complete", nil
}

func isValidTime(timeRange []int, dayRange []int) bool {
	return true
}

func checkStopCondition(timeRange []int, dayRange []int) bool {
	return false
}
