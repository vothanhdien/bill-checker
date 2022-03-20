package activity

type QueryArgs struct {
	CusCode string
}

type QueryResult struct {
	IsHasBill bool
}

type Checker interface {
	Check(cusCode string) (bool, error)
}

type BillChecker struct {
	Checker Checker
}

func (bc *BillChecker) QueryBill(input *QueryArgs) (*QueryResult, error) {
	b, err := bc.Checker.Check(input.CusCode)
	if err != nil {
		return nil, err
	}
	return &QueryResult{
		IsHasBill: b,
	}, nil
}
