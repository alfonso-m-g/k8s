# Deployment

### Create a docker container

`docker build -t kubernetes-frontend-nginx -f Dockerfile .`

### Create the deployment and the NodePort service with the yaml file

`kubectl apply -f frontend.yaml`

### Get the port where you can access the frontend with your computers IP address

`kubectl get services`