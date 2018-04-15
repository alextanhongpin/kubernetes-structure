# The IP of the container you want to test
IP := 30420
NAMESPACE := default

start:
	@minikube start --extra-config=apiserver.Authorization.Mode=RBAC
	@eval $(minikube docker-env)

stop:
	@minikube stop

dashboard:
	@minikube dashboard

service:
	@kubectl apply -f services/${NAME} --record

test:
	@curl http://$(shell minikube ip):${IP}

# Switch to a different namespace, so that you don't have to specify the namespace all the time
# make switch NAMESPACE=default
context:
	@kubectl config set-context $(shell kubectl config current-context) --namespace=${NS}

getcontext:
	@kubectl config get-contexts

l5d-verify:
	@kubectl -n ${NAMESPACE} port-forward $(shell kubectl -n ${NAMESPACE} get pod -l app=l5d -o jsonpath='{.items[0].metadata.name}') 9990 &

OUTGOING_PORT := $(shell kubectl get svc l5d -o jsonpath='{.spec.ports[0].nodePort}' -n ${NAMESPACE})
L5D_INGRESS_LB := http://$(shell minikube ip):${OUTGOING_PORT}

l5d-test:
	@curl -H "Host: hello" ${L5D_INGRESS_LB}
	http_proxy=${L5D_INGRESS_LB} curl -s http://hello
	http_proxy=${L5D_INGRESS_LB} curl -s http://world

ADMIN_PORT=$(shell kubectl get svc l5d -o jsonpath='{.spec.ports[?(@.name=="admin")].nodePort}')
l5d-dashboard:
	@open http://$(shell minikube ip):${ADMIN_PORT}

conduit-test:
	@curl http://$(shell minikube ip):$(shell kubectl get svc echo -o jsonpath='{.spec.ports[0].nodePort}')

conduit-simulate:
	wrk -d300 http://$(shell minikube ip):$(shell kubectl get svc echo -o jsonpath='{.spec.ports[0].nodePort}')