package activity

import "time"

type QueryArgs struct {
	CusCode string
}

type QueryResult struct {
	IsHasBill bool
}

func QueryBill(input *QueryArgs) (*QueryResult, error) {
	// do stub
	time.Sleep(5 * time.Second)
	return &QueryResult{
		IsHasBill: true,
	}, nil
}
