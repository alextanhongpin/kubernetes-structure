# Linkerd-Namerd

Runs linkerd in daemonset, in linker-to-linker mode, using namerd to route requests,
and handling edge traffic on a separate linkerd router.

## Setup Namespaces

```bash
$ kubectl create ns linkerd
namespace "linkerd" created
$ kubectl create ns namerd
namespace "namerd" created
```

## Setup Roles

```bash
$ kubectl apply -f roles
```

Output:

```bash
clusterrole.rbac.authorization.k8s.io "cluster-writer" configured
clusterrole.rbac.authorization.k8s.io "cluster-reader" configured
clusterrolebinding.rbac.authorization.k8s.io "cluster-write" configured
clusterrolebinding.rbac.authorization.k8s.io "cluster-read" configured
rolebinding.rbac.authorization.k8s.io "sd-build-write" unchanged
clusterrole.rbac.authorization.k8s.io "linkerd-endpoints-reader" configured
clusterrolebinding.rbac.authorization.k8s.io "linkerd-role-binding" configured
clusterrolebinding.rbac.authorization.k8s.io "l5d-linkerd-role-binding" configured
clusterrolebinding.rbac.authorization.k8s.io "namerd-linkerd-role-binding" configured
clusterrole.rbac.authorization.k8s.io "namerd-dtab-storage" configured
clusterrolebinding.rbac.authorization.k8s.io "namerd-role-binding" configured
clusterrolebinding.rbac.authorization.k8s.io "l5d-namerd-role-binding" configured
clusterrolebinding.rbac.authorization.k8s.io "default-namerd-role-binding" configured
```

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

## Setup Services

Blue Service:

```bash
$ kubectl apply -f services/blue
deployment.apps "blue" created
service "blue" created
```

Green Service:

```bash
$ kubectl apply -f services/green
deployment.apps "green" created
service "green" created
```

## Test

### Green Service

Our namerd `internal.dtab` settings is as follow, which points all service call to the blue service:

```dtab
/srv        => /#/io.l5d.k8s/default/http;
/host       => /srv;
/tmp        => /srv;
/svc        => /host;
/svc/green  => /srv/blue;
```

```bash
$ curl -H "Host: green" localhost:4140
```

Output:

```bash
{"hostname":"blue-57997758d6-s28r7","text":"hello"}
```


### Blue Service


```bash
$ curl -H "Host: blue" localhost:4140
```

Output:

```bash
{"hostname":"blue-57997758d6-hh2lh","text":"hello"}
```

## Simulate Traffic

```bash
$ wrk -d300s -H "Host: green" http://localhost:4140
# or
$ make -f Namerd.mk simulate
```

## Shift Traffic

Shift ten percent of the traffic to green service:

```bash
$ make -f Namerd.mk tenth
```

Shift half the traffic to green service:

```bash
$ make -f Namerd.mk half
```

Shift completely to green service:

```bash
$ make -f Namerd.mk full
```

## Kill

Killing the green service will bring the service back to blue automagically :smile:. Awesome right?

## Destroy

Remove blue and green services:

```bash
$ kubectl delete --all all
```

Remove linkerd:

```bash
$ kubectl delete --all all -n linkerd
```

Remove namerd:

```bash
$ kubectl delete --all all -n namerd
```