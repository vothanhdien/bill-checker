worker:
	go run cmd/worker/main.go
check:
	go run cmd/checker/main.go


.PHONY: worker run