# replicacount

[![Go Report Card](https://goreportcard.com/badge/github.com/srinandan/replicacount)](https://goreportcard.com/report/github.com/srinandan/replicacount)

This service returns the number of read & available replicas for a given ReplicaSet

## Install

* Build the docker image
* Install the [manifest](./replicacount.yaml); Don't forget to change the `project-id` for the image URL
* Provide the pod permissions to read kubernete data

```bash

kubectl create clusterrolebinding replicacount-view --clusterrole=view --serviceaccount={namespace}:default
```

NOTE: This command gives the default role the permission.  