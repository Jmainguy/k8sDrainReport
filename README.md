# k8sDrainReport
This command is designed to point out potential issues to draining a kubernetes node.

## Build
```/bin/bash
export GO111MODULE=on
go mod init
go get k8s.io/client-go@v12.0.0
go build
```

## Usage
Login to your kubernetes cluster, then run
```/bin/bash
./k8sDrainReport
```
