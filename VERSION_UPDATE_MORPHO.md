

1. Follow CONTRIBUTING setup up to cloning the repo

2. Create a new branch for the new kubernetes version

3. Search for the old version (e.g. "1.21") and update it to the newest version supported by EKS (https://docs.aws.amazon.com/eks/latest/userguide/kubernetes-versions.html)
    - InitialClusterVersion in cluster_gcp.go (gcp not used so doesnt matter?)
    - cluster-autoscaler version in images/cluster-autoscaler/Dockerfile
    - version in generate_eks

4. Fix manager/manifests/cluster-autoscaler.yaml.j2 if needed, see yaml in https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/cloudprovider/aws auto discover examples

5. Update minor version (e.g. 0.31.2 -> 0.31.3) everywhere

6. Reset go deps with the new k8s version (probably backwards compatible, check k8s go documentation)

 - `rm -rf go.mod go.sum && go mod init gitlab.com/g-ogawa/cortex && go clean -modcache`
 - `go get k8s.io/client-go@v0.23.2 && go get k8s.io/apimachinery@v0.23.2 && go get k8s.io/api@v0.23.2`
 - `go get istio.io/client-go@1.7.3 && go get istio.io/api@1.7.3`
 - `go get github.com/aws/amazon-vpc-cni-k8s/pkg/awsutils@v1.7.1`
 - `go get gitlab.com/g-ogawa/yaml@581aea36a2e4db10f8696587e48cac5248d64f4d`
 - `go get gitlab.com/g-ogawa/go-input@8b67a7a7b28d1c45f5c588171b3b50148462b247`
 - `echo -e '\nreplace github.com/docker/docker => github.com/docker/engine v19.03.12' >> go.mod`
 - `go get -u github.com/docker/distribution`
 - `go mod tidy`
 - `make test`

 - Fix all errors that might come from updating k8s versions and dependencies (good luck)


7. Build tools with `make tools`

8. Follow Building in CONTRIBUTING to build the cli and make/push the images

9. Zip the client (`cortex/bin/cortex`), create a new folder in the s3 bucket `for-mor-cortex` with the new version  (e.g. `for-mor-cortex/cli/linux/0.31.3`) and upload the zip as `cortex.zip`.

10. Build the Python wheel in pkg/cortex/client with `python setup.py bdist_wheel` AFTER UPLOADING THE NEW CLIENT!



## Notes

- Kubernetes versions are supported for about one year by AWS so this has to be done every year
- Updating the go version is likely to break things
- Any image built based on a given cortex version needs to be rebuilt with the new version
- Running clusters cant be updated, new clusters have to be created using the new cortex version