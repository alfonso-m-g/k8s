# Commands

## Create a configmap from a file
`kubectl create configmap nginx-config --from-file examples`

# Environment variables and volumes from a config map

### First get the list of the pods
`kubectl get po`

### Example output
```
NAME                                         READY   STATUS    RESTARTS   AGE
configmap-deployment-vars-67c7c999f7-5p292   1/1     Running   0          6s
```

### Access the pods terminal
`kubectl exec -ti configmap-deployment-vars-67c7c999f7-5p292 -- sh`

### Verify that the file was written into the folder 
`cat /etc/nginx/conf.d/nginx.conf`

### Example output
```
server {
    listen       80;
    server_name  localhost;

    location / {
        root   /usr/share/nginx/html;
        index  index.html index.htm;
    }

    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }
}
```

### Verify that the environment variables were written into the pod

`echo $DB_USER`

`echo DB_HOST`
