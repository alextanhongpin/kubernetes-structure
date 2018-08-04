## Expose Port

```
# :3000 refers to the pod port
$ k port-forward deployment/k8s-node :3000
```

Output:

```
Forwarding from 127.0.0.1:64188 -> 3000
Forwarding from [::1]:64188 -> 3000
Handling connection for 64188
```

Call:

```
$ curl http://127.0.0.1:64188
```

Output:

```json
{"message":"hello world"}
```

## Alternative
But does not work...
```bash
$ curl $(minikube service k8s-node --url)

# Same as
$ curl http://$(minikube ip):<nodePort>
# Get <nodePort> from k describe <service_name>
```