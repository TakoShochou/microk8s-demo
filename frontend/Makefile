PHONEY: clean dockerize microk8s.up microk8s.down demo

clean:
	rm -f frontend.tar

dockerize:
	docker build -t frontend:local --no-cache .

microk8s.up: dockerize
	- $(MAKE) microk8s.down
	docker save frontend -o frontend.tar
	microk8s ctr image import frontend.tar
	microk8s kubectl apply -f manifest_frontend.yaml

microk8s.down:
	microk8s kubectl delete -f manifest_frontend.yaml

demo: dockerize
	docker run -it -p "8080:80" --rm --name frontend-demo frontend:local
