# monolith-auth
Example of how to enable monolith authenticated traffic into micro-services

Run monolith locally:
```sh
docker-compose -f monolith/docker-compose.yml up --build server

curl http://localhost:9292/
```

Build docker image:
```sh
docker build --tag monolith-auth:0.1.0 ./monolith
```


Run application in minikube
```
minikube start --cpus 6 --memory 6144

eval $(minikube docker-env)
docker build --tag monolith-auth:0.1.0 ./monolith
docker images | grep monolith

kubectl apply -f monolith.yaml

minikube service monolith --url
curl http://127.0.0.1:53364 -H "X-Custom-Header: 123"

kubectl delete -f monolith.yaml
```


```
curl -L https://github.com/istio/istio/releases/download/1.12.0/istio-1.12.0-osx.tar.gz > istio-1.12.0-osx.tar.gz
gunzip istio-1.12.0-osx.tar.gz
tar xopf istio-1.12.0-osx.tar

export PATH=$PATH:$(pwd)/istio-1.12.0/bin
istioctl version

istioctl install
# yes

kubectl get pods -n istio-system

kubectl label namespace default istio-injection=enabled

kubectl apply -f monolith.yaml

minikube service istio-ingressgateway -n istio-system --url
http://127.0.0.1:53555/hello

```


