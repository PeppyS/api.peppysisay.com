build-image:
	docker build -t xpeppy/api.peppysisay.com .
run-image:
	docker run -t -i --env-file .env -p 8080:8080 xpeppy/api.peppysisay.com
push-image:
	docker push xpeppy/api.peppysisay.com
start:
	go run cmd/api/api.go
