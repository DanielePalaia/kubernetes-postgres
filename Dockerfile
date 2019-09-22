# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/kubernetes-postgres

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq
RUN go get github.com/swaggo/http-swagger
RUN go get github.com/alecthomas/template

ADD conf /go
WORKDIR /go/src/kubernetes-postgres
RUN go build 

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/src/kubernetes-postgres/kubernetes-postgres

# Document that the service listens on port 8080.
EXPOSE 8080
