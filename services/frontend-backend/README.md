# Frontend calling Backend

## Quickstart

```bash
# Deploy all application.
$ k apply -f .

# Check the port of the frontend app.
$ k get service frontend --watch

# Call the frontend that calls the backend.
$ curl localhost:<PORT>

# Delete all application.
$ k delete -f .
```

Output:
```
{"message": "hello"}
```

References:
- https://kubernetes.io/docs/tasks/access-application-cluster/connecting-frontend-backend/
