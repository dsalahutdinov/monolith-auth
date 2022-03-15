FROM golang:latest


WORKDIR $GOPATH/src/github.com/dsalahutdinov/monolith-auth/favorites

COPY . .
RUN go mod init favorites

RUN go get -d -v ./...

RUN go install -v ./...
RUN go build main.go
RUN ls -l main

EXPOSE 8383

CMD ["./main"]
