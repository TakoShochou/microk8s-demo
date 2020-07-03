PHONY: build dockerize microk8s.up microk8s.down demo

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o "server.exe"

clean:
	rm -f server.exe my_server.tar

dockerize: build
	docker build -t my_server:local --no-cache .
	make clean

microk8s.up: dockerize
	- make microk8s.down
	docker save my_server -o my_server.tar
	microk8s ctr image import my_server.tar
	microk8s kubectl apply -f manifest_server.yaml

microk8s.down:
	microk8s kubectl delete -f manifest_server.yaml

demo: dockerize
	docker run -it -p "3000:3000" --rm --name demo_my_server my_server:local
