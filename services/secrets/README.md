# Kubernetes secret

This example demonstrates how to use Kubernetes secrets and accessing them through two different ways, either:

- volume mount
- environment variables.

## Quick Start

Secrets needs to base-64 encoded:

```bash
$ echo -n 'your secret' | base64
```

To decode the base64 string:
```bash
$ echo 'your base64 string' | base64 --decode
```

Steps:

```bash
# To build the docker image used for this example.
$ make docker

# To deploy all services to kubernetes
$ k apply -f .

# To destroy all services that is running in kubernetes.
$ k delete -f .
```

Validate that the secret is deployed:

```bash
$ k get secrets
```

Output:

```
NAME                  TYPE                                  DATA   AGE
default-token-j4pl9   kubernetes.io/service-account-token   3      32h
mysecret              Opaque                                2      4m37s
```

To check if the secret is pass in:

```bash
$ k get service go-secret-service
```

Output:
```bash
NAME                TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
go-secret-service   NodePort   10.102.2.173   <none>        80:31927/TCP   5m19s
```

Call the endpoint:
```bash
$ curl localhost:31927
```

Output:
```
hello world%
```

Check the logs:
```
$ k logs svc/go-secret-service
```

Output:
```
2020/09/06 18:41:46 listening to port *:8080. press ctrl + c to cancel.
2020/09/06 18:41:56 secret is "secret"
2020/09/06 18:41:56 username is "john.doe"
```
