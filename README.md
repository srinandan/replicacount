# replicacount

[![Go Report Card](https://goreportcard.com/badge/github.com/srinandan/replicacount)](https://goreportcard.com/report/github.com/srinandan/replicacount)

This service returns the number of read & available replicas for a given ReplicaSet

## Install

* Build the docker image
* Install the [manifest](./replicacount.yaml); Don't forget to change the `project-id` for the image URL
* Provide the pod permissions to read kubernete data

The service exposes a single endpoint at `/count/{replicaname}`. 

```bash

kubectl create clusterrolebinding replicacount-view --clusterrole=view --serviceaccount={namespace}:default
```

### Sample

```bash

curl http://replicacount.apigee.svc.cluster.local:8080/count/testapp
```

Output:

```bash

< HTTP/1.1 200 OK
< Content-Type: text/plain; charset=UTF-8
< Date: Tue, 10 Dec 2019 07:34:57 GMT
< Content-Length: 57
<
{"replicas":1,"ready_replicas":1,"available_replicas":1}
```

NOTE: This command gives the default role the permission.  