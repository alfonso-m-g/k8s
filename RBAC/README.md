### Generate key
`openssl genrsa -out alfonso.key 2048`

### Generate Certificate Sign Request
`openssl req -new -key alfonso.key -out alfonso.csr -subj "/CN=alfonso/O=dev"`

### View certificate authority
`kubectl config view`

```
apiVersion: v1
clusters:
- cluster:
    certificate-authority: /home/alfonso/.minikube/ca.crt
    extensions:
    - extension:
        last-update: Tue, 17 Aug 2021 09:45:57 CDT
        provider: minikube.sigs.k8s.io
        version: v1.22.0
      name: cluster_info
    server: https://192.168.100.38:8443
  name: minikube
contexts:
- context:
    cluster: minikube
    namespace: development
    user: minikube
  name: dev-context
- context:
    cluster: minikube
    extensions:
    - extension:
        last-update: Tue, 17 Aug 2021 09:45:57 CDT
        provider: minikube.sigs.k8s.io
        version: v1.22.0
      name: context_info
    namespace: default
    user: minikube
  name: minikube
current-context: minikube
kind: Config
preferences: {}
users:
- name: minikube
  user:
    client-certificate: /home/alfonso/.minikube/profiles/minikube/client.crt
    client-key: /home/alfonso/.minikube/profiles/minikube/client.key
```
# Create a new user

### Sign the certificate
`sudo openssl x509 -req -in alfonso.csr -CA /home/alfonso/.minikube/ca.crt -CAkey /home/alfonso/.minikube/ca.key -CAcreateserial -out alfonso.crt -days 500`

### Certificate information
`openssl x509 -in alfonso.crt -noout -text`

### Create container with Certificate
`docker run --rm -ti -v $PWD:/test -w /test -v /home/alfonso/.minikube/ca.crt:/ca.crt -v /usr/bin/kubectl:/usr/bin/kubectl alpine sh`

### Configure a minikube cluster
`kubectl config set-cluster minikube --server=https://192.168.100.38:8443 --certificate-authority=/home/alfonso/.minikube/ca.crt`

### Set credentials
`kubectl config set-credentials alfonso --client-certificate=alfonso.crt --client-key=alfonso.key`

### View configuration
`kubectl config view`

### Get the current context
`kubectl config current-context`

### Set context
`kubectl config set-context alfonso --cluster=minikube --user=alfonso`

### Enable RBAC
`minikube start --vm-driver=none --^Ctra-config=apiserver.authorization-mode=RBAC`

### Change context
`kubectl config use-context alfonso`

### List pods on new context
`kubectl get pods`

`Error from server (Forbidden): pods is forbidden: User "alfonso" cannot list resource "pods" in API group "" in the namespace "default"`

# Roles 

### Create a role
`kubectl apply -f alfonso-pods.yaml`

### Get roles
`kubectl get roles`

### Describe roles
`kubectl describe roles`

```
Name:         pod-reader
Labels:       <none>
Annotations:  <none>
PolicyRule:
  Resources  Non-Resource URLs  Resource Names  Verbs
  ---------  -----------------  --------------  -----
  pods       []                 []              [get watch list]
```

### Describe a role binding
`kubectl describe rolebinding read-pods`

```
Name:         read-pods
Labels:       <none>
Annotations:  <none>
Role:
  Kind:  Role
  Name:  pod-reader
Subjects:
  Kind  Name     Namespace
  ----  ----     ---------
  User  alfonso  
```

# Cluster Roles

### Get cluster roles
`kubectl get clusterroles`