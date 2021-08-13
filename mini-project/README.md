<h1>(Test) Execute the following commands to run the container and the application

`<addr>`docker run --rm -dti -v $PWD/:/go --net host --name golang golang bash

`<addr>`docker exec -it golang bash

`<addr>`go run main.go



<h1>(Deploy) 

Create a docker container

`<addr>`docker build -t kubernetes-hands-on -f Dockerfile .

Run the container

`<addr>`docker run -d -p 8080:8080 --name kubernetes-example kubernetes-go}