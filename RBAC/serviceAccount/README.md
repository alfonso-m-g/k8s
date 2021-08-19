# Service Accounts

### Execute a pod
`kubectl exec -ti podtest3 -- sh`

### List the pods from the kubernetes API
`curl http://kubernetes/api/v1/namespaces/default/pods`

```
{
  "kind": "Status",
  "apiVersion": "v1",
  "metadata": {
    
  },
  "status": "Failure",
  "message": "pods is forbidden: User \"system:anonymous\" cannot list resource \"pods\" in API group \"\" in the namespace \"default\"",
  "reason": "Forbidden",
  "details": {
    "kind": "pods"
  },
  "code": 403
}
```

### Get token and save it in an environment variable
`export TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)`

### Fetch from API with authentication token
`curl -H "Authorization: Bearer ${TOKEN}" https://kubernetes/api/v1 --insecure`

### Fetch pods from API
`curl -H "Authorization: Bearer ${TOKEN}" https://kubernetes/api/v1/namespaces/default/pods --insecure`

### Fetch deployments from API
`curl -H "Authorization: Bearer ${TOKEN}" https://kubernetes/apis/apps/v1/namespaces/default/deployments --insecure`