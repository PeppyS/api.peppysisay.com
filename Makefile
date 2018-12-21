build-image:
	docker build -t api .
run-image:
	docker run -t -i --env-file .env -p 8080:8080 api
start:
	go run cmd/api/api.go
