build:
	docker build -t go-live .

run-worker:
	docker run -it --network host go-live worker

run-server1:
	docker run -it --network host go-live server1

run-server2:
	docker run -it --network host go-live server1

lint:
	golangci-lint run -v
	@echo "DONE ✅"

clean:
	rm -r output/*