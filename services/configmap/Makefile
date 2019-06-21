docker:
	@docker build -t alextanhongpin/go-server .

run:
	@docker run alextanhongpin/go-server

# To call the service through port-forwarding
# kubectl port-forward deployment/go-server-deployment 8080:8080
# # Listens to port 3000 on host, which maps to port 8080 in container.
# kubectl port-forward service/go-server-service 3000:8080
# NOTE: calling through NodePort or LoadBalancer http://localhost:3xxxx (nodePort) does not work.
