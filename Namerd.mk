
setup:
	$(eval NAMERD_IP=$(shell kubectl get svc namerd -o jsonpath='{.spec.ports[?(@.name=="http")].nodePort}' -n namerd))
	@curl -v -XPUT -d @services/namerd-http/internal.dtab -H "Content-Type: application/dtab" http://${KUBE_IP}:${NAMERD_IP}/api/1/dtabs/internal
	@curl -v -XPUT -d @services/namerd-http/external.dtab -H "Content-Type: application/dtab" http://${KUBE_IP}:${NAMERD_IP}/api/1/dtabs/external

original:
	$(eval NAMERD_IP=$(shell kubectl get svc namerd -o jsonpath='{.spec.ports[?(@.name=="http")].nodePort}'))
	@curl -v -XPUT -d @services/namerd-http/internal.dtab -H "Content-Type: application/dtab" http://${KUBE_IP}:${NAMERD_IP}/api/1/dtabs/internal

call-blue:
	$(eval LINKERD_IP = $(shell kubectl get svc l5d -o jsonpath='{.spec.ports[?(@.name=="outgoing")].nodePort}'))
	curl -H "Host: blue" http://${KUBE_IP}:${LINKERD_IP}

call-green:
	$(eval LINKERD_IP = $(shell kubectl get svc l5d -o jsonpath='{.spec.ports[?(@.name=="outgoing")].nodePort}'))
	curl -H "Host: green" http://${KUBE_IP}:${LINKERD_IP}

simulate:
	# $(eval LINKERD_IP = $(shell kubectl get svc l5d -o jsonpath='{.spec.ports[?(@.name=="outgoing")].nodePort}'))
	# wrk -d300s -H "Host: green" http://${KUBE_IP}:${LINKERD_IP}
	wrk -d300s -H "Host: green" http://localhost:4140

tenth:
	$(eval NAMERD_IP=$(shell kubectl get svc namerd -o jsonpath='{.spec.ports[?(@.name=="http")].nodePort}' -n namerd))
	@curl -v -XPUT -d @services/namerd-http/tenth.dtab -H "Content-Type: application/dtab" http://${KUBE_IP}:${NAMERD_IP}/api/1/dtabs/internal

half:
	$(eval NAMERD_IP=$(shell kubectl get svc namerd -o jsonpath='{.spec.ports[?(@.name=="http")].nodePort}' -n namerd))
	@curl -v -XPUT -d @services/namerd-http/half.dtab -H "Content-Type: application/dtab" http://${KUBE_IP}:${NAMERD_IP}/api/1/dtabs/internal


full:
	$(eval NAMERD_IP=$(shell kubectl get svc namerd -o jsonpath='{.spec.ports[?(@.name=="http")].nodePort}' -n namerd))
	@curl -v -XPUT -d @services/namerd-http/full.dtab -H "Content-Type: application/dtab" http://${KUBE_IP}:${NAMERD_IP}/api/1/dtabs/internal
