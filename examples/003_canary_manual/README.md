# Canary Deployment


We create a simple server running on port 8080.


We have two separate deployments, which will be served through a single service.



However, we can't port forward the service, because the request won't be load balanced. Only one pod will be port-forwarded.

We can instead ssh into one of the pods, the make a request to the cluster ip:

```bash
$ k exec -it <pod-name> -- sh
$ wget -qO- <cluster-ip>
```

The cluster ip can be obtained by running:

```bash
$ k get svc
NAME         TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
kubernetes   ClusterIP      10.96.0.1       <none>        443/TCP        291d
my-service   LoadBalancer   10.106.239.21   <pending>     80:30127/TCP   3m16s
```



When doing:
```
$ wget -qO- 10.106.239.21
```

We expect the requests to be mixed between canary and stable.


Note that we tried installing metallb, however, we still cannot access the service.
