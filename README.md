# monolith-auth
Example of how to enable monolith authenticated traffic into micro-services

Run monolith locally:
```sh
docker-compose -f monolith/docker-compose.yml up --build server

curl http://localhost:9292/
```

Build docker image:
```sh
docker build --tag monlith-auth:0.1.3 ./monolith
```
