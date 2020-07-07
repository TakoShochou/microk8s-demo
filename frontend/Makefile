PHONEY: clean dockerize microk8s.up microk8s.down demo

clean:
	- rm -f frontend.tar
	- docker rmi frontend frontend.slim

dockerize:
	docker build -t frontend --no-cache .
	docker run --rm -v /var/run/docker.sock:/var/run/docker.sock dslim/docker-slim build frontend
	docker rmi frontend

microk8s.up:
	- $(MAKE) microk8s.down
	docker save frontend.slim -o frontend.slim.tar
	microk8s ctr image import frontend.slim.tar
	microk8s kubectl apply -f manifest_frontend.yaml

microk8s.down:
	microk8s kubectl delete -f manifest_frontend.yaml

demo: dockerize
	docker run -it -p "8080:80" --rm --name frontend-demo frontend
