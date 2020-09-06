# Reading environment variables

This example demonstrates how to read environment variables for Kubernetes by specifying it directly in the `env` or reading it from `ConfigMap`.


## Quickstart

```bash
# Build the docker image required for this example.
$ make docker

# Deploy the service.
$ k apply -f .

# Find out the node port.
$ k get service hello-env-service

# To stop all services.
$ k delete -f .
```

Output:
```bash
NAME                TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
hello-env-service   NodePort   10.100.239.215   <none>        80:31963/TCP   110s
```

Call the service:
```bash
$ curl localhost:31963
```

Output:
```bash
{"env": "this is a greeting", "msg", "what is the message?"}%
```
