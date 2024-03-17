# Canary Ingress

In this demo, we use an Ingress to route traffic to two separate deployments, which will be served through a single service.

## Setup

First, install the `nginx-ingress` controller following the examples in the folder `examples/002_nginx_ingress`.

## Deploy

```bash
$ k apply -f .
$ k port-forward svc/my-service 3000:80

$ curl localhost:3000/hello
Hello, World! - stable%
$ curl localhost:3000/hello
Hello, World! - canary%
$ curl localhost:3000/hello
Hello, World! - canary%
$ curl localhost:3000/hello
Hello, World! - canary%
$ curl localhost:3000/hello
Hello, World! - stable%
```
