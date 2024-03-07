build:
	docker build -t go-live

run: 
	docker run -d go-live -p 8000:8000