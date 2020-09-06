## Service-to-service communication

This example demonstrates how service-to-service communication in Kubernetes works. We have the following server:

- server: a single endpoint that returns hello world when called
- client: a single endpoint that calls the server endpoint and returns the server response

## Quickstart

```bash
# Build the docker images required for this example.
$ make docker-client
$ make docker-server

# Deploy the client and server service/deployment to Kubernetes
$ k apply -f .

# Drop the services.
$ k delete -f .
```

To check the endpoint:
```bash
$ k get client-service
```
Output:

```
> k get service client-service
zsh: correct 'service' to 'services' [nyae]? n
NAME             TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
client-service   NodePort   10.111.71.115   <none>        80:30025/TCP   4m8s
```

Call the endpoint:
```bash
$ curl localhost:30025
```

Output:
```bash
$ curl localhost:30025
{"hello": "world"}%
```

Check the server logs to see that it has been called:

```bash
$ k logs service/server-service
2020/09/06 13:29:16 listening to port *:8080. press ctrl + c to cancel.
calling server
```
