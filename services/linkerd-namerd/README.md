# Linkerd-Namerd

Runs linkerd in daemonset, in linker-to-linker mode, using namerd to route requests,
and handling edge traffic on a separate linkerd router.

Used in conjunction with `services/namerd-http` and blablabla.


## Call 

curl -H "Host: hello" http://$(minikube ip):$(kubectl get svc l5d -o jsonpath='{.spec.ports[?(@.name=="outgoing")].nodePort}')

Hello (172.17.0.15) world (172.17.0.19)!!%

# World
curl -H "Host: world" http://$(minikube ip):$(kubectl get svc l5d -o jsonpath='{.spec.ports[?(@.name=="outgoing")].nodePort}')
world (172.17.0.18)!%