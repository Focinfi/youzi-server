FROM golang:1.6.2
ADD . $GOPATH/src/github.com/focinfi/youzi-server/
RUN apt-get update && apt-get install -y git
RUN go get github.com/focinfi/youzi-server
RUN CGO_ENABLED=0 go install -a github.com/focinfi/youzi-server
WORKDIR $GOPATH/src/github.com/focinfi/youzi-server/
EXPOSE 1234
CMD ["youzi-server"]
