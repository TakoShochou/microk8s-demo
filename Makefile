.PHONY: clean dockerize microk8s.up microk8s.down

dockerize:
	$(MAKE) -C ./demo-server1 dockerize
	$(MAKE) -C ./frontend dockerize

clean:
	$(MAKE) -C ./demo-server1 clean
	$(MAKE) -C ./frontend clean

microk8s.up:
	$(MAKE) -C ./demo-server1 microk8s.up
	#$(MAKE) -C ./demo-server2 microk8s.up
	$(MAKE) -C ./frontend microk8s.up

microk8s.down:
	- $(MAKE) -C ./demo-server1 microk8s.down
	- $(MAKE) -C ./demo-server2 microk8s.down
	- $(MAKE) -C ./frontend microk8s.down
