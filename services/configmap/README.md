## Calling the service through docker for mac

### Quickstart

```bash
# Build the docker image
$ make docker

# Deploy the application to kubernetes
$ alias k=kubectl
$ k apply -f .

# To destroy the application.
$ k delete -f .
```

https://github.com/docker/for-mac/issues/2445

I encountered this issue recently, and surprisingly it works on one machine but not on another. Here's what I observed on my working machine:

1) It doesn't have minikube installed, so the only context is docker-for-desktop. So I uninstalled minikube on the other machine using the steps https://github.com/kubernetes/minikube/issues/1043#issuecomment-382031886.
2) It doesn't work, but at least I cleared minikube. The other observation is that the status of Kubernetes on the machine that is not working is showing _Kubernetes is starting_ permanently. I tried to resolve that through https://github.com/docker/for-mac/issues/2990. The only solution that worked for me is to _reset to factory defaults_, since _restart_, _reset kubernetes cluster_, _reset disk image_ options are not working.

The status shown now is _Kubernetes is running_. I tried deploying my service and I can now call it. This is how my service definition looks like:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: go-server-service
spec:
  type: NodePort
  ports:
  - protocol: TCP
    port: 8080
  selector:
    app: go-server
```

Here's what I see with the command `k get svc` (note that I use the `alias k=kubectl`:
```bash
$ k get svc
NAME                TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
go-server-service   NodePort    10.99.164.31   <none>        8080:31422/TCP   8m
kubernetes          ClusterIP   10.96.0.1      <none>        443/TCP        11m
```

Here's the output when I make a `curl`:
```bash
$ curl localhost:31422
{"message": "hello world"}
```
