build:
	docker build -t go-live .

run-worker:
	docker run -ti --rm --network host -v OutputVolume:/app go-live worker

run-server1:
	docker run -ti --rm --network host -v OutputVolume:/app go-live server1

run-server2:
	docker run -ti --rm --network host -v OutputVolume:/app go-live server2

run-local-worker:
	go run main.go worker

run-local-server1:
	go run main.go server1

run-local-server2:
	go run main.go server2

lint:
	golangci-lint run -v
	@echo "DONE ✅"

clean:
	rm -r output/*