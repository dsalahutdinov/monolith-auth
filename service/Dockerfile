FROM golang:latest 


WORKDIR $GOPATH/src/github.com/dsalahutdinov/monolith-auth/service

COPY . .
RUN go mod init service

RUN go get -d -v ./...

RUN go install -v ./...
RUN go build main.go
RUN ls -l main

EXPOSE 8383

CMD ["./main"]

