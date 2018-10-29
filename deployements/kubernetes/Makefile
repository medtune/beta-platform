# build server container
build:
	GOOS=linux go build -o server
	@eval $$(minikube docker-env) ;\
	docker build -t test-app:v1 .

# create kubernetes cluster & service

create:
	kubectl create -f deploy/k8s.yaml

# delete kubenetes cluster & service
delete:
	kubectl delete -f deploy/k8s.yaml

# clean working directory
clean:
	rm -f server

# create heapster, graphana & influxdb 
# monitoring service
heapster-up:
	kubectl create -f heapster/

# delete heapster cluster
heapster-down:
	kubectl delete -f heapster/

# show test app
show:
	minikube service test-svc

# app logs
logs:
	kubectl logs -l app=my-app

# list minikube services for urls
list-svc:
	minikube service list