build:
	docker build -t go-live .

run-worker:
	docker run -ti --rm --network host -v OutputVolume:/app go-live worker

run-server1:
	docker run -ti --rm --network host -v OutputVolume:/app go-live server1

run-server2:
	docker run -ti --rm --network host -v OutputVolume:/app go-live server2

lint:
	golangci-lint run -v
	@echo "DONE ✅"

clean:
	rm -r output/*