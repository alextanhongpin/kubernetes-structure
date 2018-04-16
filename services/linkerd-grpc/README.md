# Linkerd gRPC

## Create Linkerd

```bash
$ kubectl apply -f services/linkerd-grpc
```

## Create gRPC Server and Client

```bash
$ kubectl apply -f services/grpc
```

Note that the client hostIP has to be obtained through the environment variable:

```yml
        env:
        - name: HOST_IP
          # curl localhost:8080/api/v1/namespaces/default/pods/l5d-h8t8z | grep hostIP
          # value: 192.168.99.100:4140
          valueFrom:
            fieldRef:
              fieldPath: status.hostIP
        - name: PORT
          value: $(HOST_IP):4140
```

## List Pods

```bash
NAME                           READY     STATUS             RESTARTS   AGE
grpc-client-6fb794745c-fb6vt   0/1       Completed     0          3s
grpc-server-b4946ffcc-jpgbt    1/1       Running       0          25m
l5d-7r4sl                      2/2       Running       0          18m
```

## Admin UI

```bash
$ kubectl port-forward l5d-7r4sl 9990:9990
Forwarding from 127.0.0.1:9990 -> 9990
Forwarding from [::1]:9990 -> 9990
```

## Validate

Our gRPC client is not a long-living service - it will only make a call to the gRPC server endpoint once and terminates.

To validate that the client is calling the server, we will get the pod logs:

```bash
$ kubectl logs grpc-client-6fb794745c-fb6vt
greeting from grpc-server-b4946ffcc-jpgbt: Hello, John Doe

$ kubectl logs grpc-server-b4946ffcc-jpgbt
got metadata: secret {"_internal_repr":{"authorization":["secret"],"user-agent":["grpc-node/1.10.1 grpc-c/6.0.0-pre1 (linux; chttp2; glamorous)"],"l5d-dst-service":["/svc/echo.Echo"],"via":["h2 linkerd, h2 linkerd"],"l5d-dst-client":["/%/io.l5d.k8s.localnode/172.17.0.4/#/io.l5d.k8s/default/grpc/grpc-server"],"l5d-dst-residual":["/echo"],"l5d-ctx-trace":["PGG8OkSOXzU6B08pEDCSWZnOL7CwIRaxAAAAAAAAAAA="],"l5d-reqid":["99ce2fb0b02116b1"]}}
```

## Load-balancing

We can scale the gRPC server replicas to `3` and trigger the client multiple times. It should call the different gRPC server hostname:

```bash
greeting from grpc-server-b4946ffcc-jpgbt: Hello, John Doe
greeting from grpc-server-b4946ffcc-6mvj6: Hello, John Doe
```