package workflow

import (
	"time"

	"github.com/vothanhdien/bill-checker/internal/activity"
	"go.temporal.io/sdk/workflow"
)

type CheckBillWFInput struct {
	CusCode string
}

type CheckBillWFAttempt struct {
	T time.Time
	B bool
}

type CheckBillWFOutput struct {
	Attempts []CheckBillWFAttempt
}

func CheckBillWorkFlow(ctx workflow.Context, in *CheckBillWFInput) (*CheckBillWFOutput, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)
	var result activity.QueryResult
	err := workflow.ExecuteActivity(ctx, "QueryBill", activity.QueryArgs{CusCode: in.CusCode}).Get(ctx, &result)
	return &CheckBillWFOutput{
		Attempts: []CheckBillWFAttempt{
			{
				T: time.Now(),
				B: result.IsHasBill,
			},
		},
	}, err
}
