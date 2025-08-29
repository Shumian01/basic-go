.PHONY: docker
docker:
	@set GOOS=linux
	@set GOARCH=amd64
	@go build -tags=k8s -o webook .
	@docker rmi -f shumian/webook:v0.0.1
	@docker build -t shumian/webook:v0.0.1 .
kubectl:
	@kubectl delete deployment webook
	@kubectl apply -f k8s-webook-deployment.yaml