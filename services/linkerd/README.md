kubectl logs l5d-wrwxp l5d -n linkerd
k apply -f roles/
curl -H "Host: hello.linkerd" http://$(minikube ip):32056