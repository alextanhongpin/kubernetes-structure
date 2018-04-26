# Get the http api port

## Execute namerd

Setup internal linker:

```bash
$ curl -v -XPUT -d @services/namerd-http/internal.dtab -H "Content-Type: application/dtab" http://$(minikube ip):$(kubectl get svc namerd -o jsonpath='{.spec.ports[?(@.name=="http")].nodePort}')/api/1/dtabs/internal
```

Setup external linker:

```bash
$ curl -v -XPUT -d @services/namerd-http/external.dtab -H "Content-Type: application/dtab" http://$(minikube ip):$(kubectl get svc namerd -o jsonpath='{.spec.ports[?(@.name=="http")].nodePort}')/api/1/dtabs/external
```

## Logs

```bash
$ kubectl logs $(kubectl get po -l app=namerd -o jsonpath='{.items[0].metadata.name}') namerd
```

## View ui

```bash
$ kubectl port-forward $(kubectl get po -l app=namerd -o jsonpath='{.items[0].metadata.name}') 9991
```
