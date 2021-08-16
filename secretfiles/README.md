### Create a secret
`kubectl create secret generic mysecret --from-file secrets.txt`

# Create environment variables

`export USER=alfonso`

`export PASSWORD=pachi`

### Pass the variables to a temporal yaml file

`envsubst < secure.yaml > temp.yaml`

### Create the secret

`kubectl apply -f temp.yaml`