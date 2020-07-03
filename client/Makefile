PHONEY: serve dockerize

clean:
	rm -f my_client.tar

dockerize:
	docker build -t my_client:local --no-cache .

microk8s.up: dockerize
	- make microk8s.down
	docker save my_client -o my_client.tar
	microk8s ctr image import my_client.tar
	microk8s kubectl apply -f manifest_client.yaml

microk8s.down:
	microk8s kubectl delete -f manifest_client.yaml

demo: dockerize
	docker run -it -p "8080:80" --rm --name demo_my_client my_client:local
