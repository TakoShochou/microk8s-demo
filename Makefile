.PHONY: microk8s.up microk8s.down

microk8s.up:
	$(MAKE) -C ./demo-server1 microk8s.up
	#$(MAKE) -C ./demo-server2 microk8s.up
	$(MAKE) -C ./frontend microk8s.up

microk8s.down:
	- $(MAKE) -C ./demo-server1 microk8s.down
	- $(MAKE) -C ./demo-server2 microk8s.down
	- $(MAKE) -C ./frontend microk8s.down
