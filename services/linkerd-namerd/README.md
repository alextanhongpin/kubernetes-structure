# Linkerd-Namerd

Runs linkerd in daemonset, in linker-to-linker mode, using namerd to route requests,
and handling edge traffic on a separate linkerd router.

Used in conjunction with `services/namerd-http` and blablabla.

## Run Linkerd and Namerd

Note that the commands are all invoked from the `default` context.

```bash
$ kubectl apply -f services/linkerd-namerd -n linkerd
$ kubectl apply -f services/namerd-http -n namerd
```

## Open UI for Linkerd and Namerd

```bash
$ kubectl port-forward svc/l5d 9990:9990 -n linkerd
$ kubectl port-forward svc/namerd 9991:9991 -n namerd
```

## Expose the Linkerd HTTP Port

```bash
$ kubectl port-forward svc/l5d 4140:4140 -n linkerd
```

## Setup Namerd Dtabs

TODO: Automate it with `Job`.

```bash
$ make -f Namerd.mk setup
```

## Call 

curl -H "Host: hello" http://$(minikube ip):$(kubectl get svc l5d -o jsonpath='{.spec.ports[?(@.name=="outgoing")].nodePort}' -n l5d)

Hello (172.17.0.15) world (172.17.0.19)!!%

# World
curl -H "Host: world" http://$(minikube ip):$(kubectl get svc l5d -o jsonpath='{.spec.ports[?(@.name=="outgoing")].nodePort}' -n l5d)
world (172.17.0.18)!%

curl -H "Host: blue" http://$(minikube ip):$(kubectl get svc l5d -o jsonpath='{.spec.ports[?(@.name=="outgoing")].nodePort}' -n l5d)