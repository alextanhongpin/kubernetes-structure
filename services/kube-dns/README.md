## Minikube

```bash
$ minikube addons list
```

Shows that `kube-dns` is enabled:

```
- addon-manager: enabled
- coredns: disabled
- dashboard: enabled
- default-storageclass: enabled
- efk: disabled
- freshpod: disabled
- heapster: disabled
- ingress: disabled
- kube-dns: enabled
- metrics-server: disabled
- nvidia-driver-installer: disabled
- nvidia-gpu-device-plugin: disabled
- registry: disabled
- registry-creds: disabled
- storage-provisioner: enabled
```


But, the pod does not exist when running `k get po -n kube-system`:

```
NAME                                    READY     STATUS             RESTARTS   AGE
etcd-minikube                           1/1       Running            0          6m
kube-addon-manager-minikube             1/1       Running            4          3h
kube-apiserver-minikube                 1/1       Running            0          6m
kube-controller-manager-minikube        1/1       Running            0          6m
kube-scheduler-minikube                 1/1       Running            4          3h
kubernetes-dashboard-5498ccf677-z2knl   0/1       CrashLoopBackOff   20         3h
storage-provisioner                     0/1       CrashLoopBackOff   20         3h
```

## Solution

Run `kube-dns` it separately.

https://stackoverflow.com/questions/50680652/minikube-kube-dns-addon-not-worki
ng-cant-enable-it