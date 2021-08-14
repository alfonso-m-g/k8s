# Useful commands

## Get the description of the limitrange
`kubectl describe limitranges mem-limit-range -n dev`

### Example output
```
Name:       mem-limit-range
Namespace:  dev
Type        Resource  Min  Max  Default Request  Default Limit  Max Limit/Request Ratio
----        --------  ---  ---  ---------------  -------------  -----------------------
Container   cpu       -    -    500m             1              -
Container   memory    -    -    256Mi            512Mi          -
```

## Get the description of the limit ranges of a pod
`kubectl get pod podx -o yaml -n dev | grep -i limits -C3`

### Example output
```
    imagePullPolicy: IfNotPresent
    name: container1
    resources:
      limits:
        cpu: "1"
        memory: 512Mi
      requests:
```