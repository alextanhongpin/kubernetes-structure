# k8s-node

Sample docker build for nodejs express.

For deployment to docker-compose, see `docker-compose.yml`. For kubernetes, refer to the service folder.


## Comparison between the bare `alpine:3.7` vs `mhart/alpine-node:10.8.0` image size

```
alextanhongpin/k8s-node-alpine     latest                  2587468792d4        About a minute ago   58.9MB
alextanhongpin/k8s-node            latest                  8503576204ad        2 days ago           81.2MB
```