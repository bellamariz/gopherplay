build-image:
	docker build -t go-live

run-image: 
	docker run -d go-live -p 8000:8000

run-local:
	VIDEO_PATH=./assets/BigBuckBunny.mp4 \
	go run main.go