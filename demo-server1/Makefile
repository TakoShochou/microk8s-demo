.PHONY: build clean dockerize microk8s.up microk8s.down demo


APP_NAME = demo-server1
REGISTRY = localhost:5000

CPU := $(shell uname -m)
ifeq ("$(CPU)", "aarch64")
ARCH = arm64
else
ARCH = amd64
endif
MANIFEST = manifest_$(APP_NAME)_$(ARCH).yaml

build:
	@echo CPU=$(CPU) ARCH=$(ARCH)
	GOOS=linux GOARCH=$(ARCH) CGO_ENABLED=0 go build -o "$(APP_NAME).exe"

clean:
	- rm -f $(APP_NAME).exe $(APP_NAME).slim.tar
	- docker rmi $(APP_NAME) $(APP_NAME).slim

dockerize: build
	docker build -t $(APP_NAME) --no-cache --build-arg EXE_FILE=$(APP_NAME).exe .
	docker tag $(APP_NAME) $(REGISTRY)/$(APP_NAME):$(ARCH)
	docker push $(REGISTRY)/$(APP_NAME):$(ARCH)

microk8s.up:
	- $(MAKE) microk8s.down
	microk8s kubectl apply -f $(MANIFEST)

microk8s.down:
	microk8s kubectl delete -f $(MANIFEST)

demo: dockerize
	docker run -it -p "30030:3000" --rm --name $(APP_NAME)-demo $(APP_NAME)
