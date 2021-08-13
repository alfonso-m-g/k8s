# Commands

## Create namespace 
`kubectl create namespace development`

## Create a new context
`kubectl config set-context dev-context --namespace=development --cluster=minikube --user=minikube`

## View context configuration
`kubectl config view`

## Switch context
`kubectl config use-context dev-context`