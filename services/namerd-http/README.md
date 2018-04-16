# Get the http api port

## Execute namerd

curl -v -XPUT -d @services/namerd-http/internal.dtab -H "Content-Type: application/dtab" http://$(minikube ip):$(kubectl get svc namerd -o jsonpath='{.spec.ports[?(@.name=="http")].nodePort}')/api/1/dtabs/internal

curl -v -XPUT -d @services/namerd-http/external.dtab -H "Content-Type: application/dtab" http://$(minikube ip):$(kubectl get svc namerd -o jsonpath='{.spec.ports[?(@.name=="http")].nodePort}')/api/1/dtabs/external

## Logs

kubectl logs $(kubectl get po -l app=namerd -o jsonpath='{.items[0].metadata.name}') namerd

## View ui

kubectl port-forward $(kubectl get po -l app=namerd -o jsonpath='{.items[0].metadata.name}') 9991

## Route the traffic by modifying the external services
