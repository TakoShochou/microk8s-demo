.PHONY: build clean dockerize microk8s.up microk8s.down demo

APP_NAME = demo-server1

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o "$(APP_NAME).exe"

clean:
	- rm -f $(APP_NAME).exe $(APP_NAME).slim.tar
	- docker rmi $(APP_NAME) $(APP_NAME).slim

dockerize: build
	docker build -t $(APP_NAME) --no-cache --build-arg EXE_FILE=$(APP_NAME).exe .
	docker run --rm -v /var/run/docker.sock:/var/run/docker.sock dslim/docker-slim build $(APP_NAME)
	docker rmi $(APP_NAME)

microk8s.up:
	- $(MAKE) microk8s.down
	docker save $(APP_NAME).slim -o $(APP_NAME).slim.tar
	microk8s ctr image import $(APP_NAME).slim.tar
	microk8s kubectl apply -f manifest_$(APP_NAME).yaml

microk8s.down:
	microk8s kubectl delete -f manifest_$(APP_NAME).yaml

demo: dockerize
	docker run -it -p "30030:3000" --rm --name $(APP_NAME)-demo $(APP_NAME)
