APP_NAME := demo-server2

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(APP_NAME).exe

.PHONY: clean
clean:
	- rm -f $(APP_NAME).exe ($APP_NAME).tar

.PHONY: demo
demo:
	$(APP_NAME).exe