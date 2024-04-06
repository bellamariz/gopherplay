build:
	docker build -t go-live .

run-worker:
	docker run -ti --rm --network host -v OutputVolume:/app go-live worker

run-server1:
	docker run -ti --rm --network host -v OutputVolume:/app go-live server1

run-server2:
	docker run -ti --rm --network host -v OutputVolume:/app go-live server2

run-discovery:
	docker run -ti --rm --network host -v OutputVolume:/app go-live discovery

run-reporter:
	docker run -ti --rm --network host -v OutputVolume:/app go-live reporter

run-origin:
	docker run -ti --rm --network host -v OutputVolume:/app go-live origin

run-local-worker:
	go run main.go worker

run-local-server1:
	go run main.go server1

run-local-server2:
	go run main.go server2

run-local-discovery:
	go run main.go discovery

run-local-reporter:
	go run main.go reporter

run-local-origin:
	go run main.go origin

lint:
	golangci-lint run -v
	@echo "DONE ✅"

clean:
	rm -r output/*