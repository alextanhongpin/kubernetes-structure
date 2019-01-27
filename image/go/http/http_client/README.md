# Kubernetes for Docker Desktop

Example of calling another service endpoint in kubernetes (http client/server example).

- run two service in kubernetes, one server and another client
- the client can call the server endpoint through http://http-server:8080/hello


## Build

Build the docker image required:

```bash
$ make build
```

## Deploy 

```bash
$ alias k=kubectl
$ k apply -f service
```

## Call the service

```bash
$ make call
```

Output:

```
curl http://localhost:31290
server
```
