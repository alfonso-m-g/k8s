(Test) Execute the following commands to run the container and the application

docker run --rm -dti -v $PWD/:/go --net host --name golang golang bash

docker exec -it golang bash

go run main.go



(Deploy) 

Create a docker container

docker build -t kubernetes-hands-on -f Dockerfile .

Run the container

docker run -d -p 8080:8080 --name kubernetes-example kubernetes-go}