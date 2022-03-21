worker:
	go run cmd/worker/main.go
check:
	go run cmd/checker/main.go
query:
	go run cmd/querier/main.go
stop:
	go run cmd/stop/main.go

.PHONY: worker run