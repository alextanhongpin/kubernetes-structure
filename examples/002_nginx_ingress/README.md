# Setting up Nginx Ingress using Helm

Installation and upgrade instructions: https://docs.nginx.com/nginx-ingress-controller/installation/installing-nic/installation-with-helm/

## Installation

We first install the nginx ingress using helm chart.

```bash
$ helm pull oci://ghcr.io/nginxinc/charts/nginx-ingress --untar --version 1.1.3

$ cd nginx-ingress

$ helm install <release-name> .
$ helm install my-release .

# Upgrade the release.
$ helm upgrade my-release .

$ helm uninstall my-release
```

## List the releases:

```bash
$ helm list -A
```


## Check ingress

```bash
alias k=kubectl
$ k get ingress
```

There's no ingress, we need to create a new ingress rule.

## Ingress Rule

What is an ingress rule? Ingress rule are a set resources that define how external traffic is routed to your services.

We want to create a new ingress rule that will route traffic to our service.

There are several ways to route traffic:
- by hostname
- by url path

https://kubernetes.github.io/ingress-nginx/examples/rewrite/

## Exposing port

How do we reach the nginx ingress?
We need to port-forward the service, which will then be accessible through localhost:3000.

```bash
k port-forward svc/my-release-nginx-ingress-controller 3000:80
```

## Reaching a service

First, create a new service that will be exposed through the ingress.

```bash
cd services/server
make build
k apply -f .
```

After applying, deploy our ingress:

```bash
$ k apply -f .
```

The service is now reachable through `localhost:3000/hello`

```bash
$ curl localhost:3000/hello
Hello, World!
```

It is not accessible through other endpoints.

```bash
$ curl localhost:3000
<html>
<head><title>404 Not Found</title></head>
<body>
<center><h1>404 Not Found</h1></center>
<hr><center>nginx/1.25.4</center>
</body>
</html>
```


The routing is as follow:

- http://localhost:3000/hello -> ingress -> k8s-service -> my-service:80
- http://localhost:3000 -> ingress (404)
