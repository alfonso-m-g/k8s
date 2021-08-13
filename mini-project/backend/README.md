# Testing

### Run the container
`docker run --rm -dti -v $PWD/:/go --net host --name golang golang bash`

## Run a bash terminal on the container
`docker exec -it golang bash`

## Run the application
`go run main.go`



# Deployment

### Create a docker container

`docker build -t kubernetes-hands-on -f Dockerfile .`

### Run the container

`docker run -d -p 8080:8080 --name kubernetes-example kubernetes-go`