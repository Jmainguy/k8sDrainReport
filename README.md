# k8sDrainReport
This command is designed to point out potential issues to draining a kubernetes node.

## Usage
Login to your kubernetes cluster, then run
```/bin/bash
./k8sDrainReport
```

## PreBuilt Binaries
Grab Binaries from [The Releases Page](https://github.com/Jmainguy/k8sDrainReport/releases)

## Build
```/bin/bash
export GO111MODULE=on
go mod init
go get k8s.io/client-go@v12.0.0
go build
```


## Examples

### Run without any cordoned/drained nodes

```
$ k8sDrainReport
Cluster: https://master.atl1.ocp.bandwidth.com:8443
There are 575 running pods in the cluster
========================================
Pods without owners
========================================
There are 0 pods without ownership

========================================
Pod Distruption Budget, Potential Issues
========================================
```
