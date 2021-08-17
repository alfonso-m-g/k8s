### View the storage classes
`kubectl get storageclass`

### Get the persistent volumes
`kubectl get pv`

### Output
```
NAME             CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                   STORAGECLASS   REASON   AGE
task-pv-volume   10Gi       RWO            Retain           Bound    default/task-pv-claim   manual                  28s
``` 

### Edit a persistent volume
`kubectl edit pv task-pv-volume`