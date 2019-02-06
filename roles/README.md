# Setting up Dashboard for Docker Desktop

https://github.com/kubernetes/dashboard

```
# Create user
$ kubectl create -f kuberneted-admin.yml

# Download image
$ kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml

# Create secure channel to kubernetes cluster
$ kubectl proxy

# Access dashboard
http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/
```
