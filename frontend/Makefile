PHONEY: clean dockerize microk8s.up microk8s.down demo

REGISTRY = localhost:5000
IMAGE = frontend

clean:
	- rm -f frontend.slim.tar
	- docker rmi $(IMAGE) $(IMAGE).slim $(REGISTRY)/$(IMAGE).slim

dockerize:
	docker build -t $(IMAGE) --no-cache .
	docker tag $(IMAGE) $(REGISTRY)/$(IMAGE)
	docker push $(REGISTRY)/frontend

microk8s.up:
	- $(MAKE) microk8s.down
	microk8s kubectl apply -f manifest_frontend.yaml

microk8s.down:
	microk8s kubectl delete -f manifest_frontend.yaml

demo: dockerize
	docker run -it -p "8080:80" --rm --name frontend-demo frontend
