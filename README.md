# dagger-tools

> The dagger engine image with a collection of tools used to be easily deployable to a Kubernetes cluster for debugging purposes.

Also inlcudes most linux-crisis-tools tools mentioned on https://www.brendangregg.com/blog/2024-03-24/linux-crisis-tools.html

## Usage

```bash
kubectl run --image mheers/dagger-tools:v0.16.3 host-dagger-tools --command sleep infinity
```

# Build

```bash
cd ci/

export $(cat .env | xargs)
dagger call build-and-push-image --registry-token=env:REGISTRY_ACCESS_TOKEN
```
