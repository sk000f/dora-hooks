FROM golang

# Copy source code to container workspace
ADD . /go/src/github.com/sk000f/metrix

# Grab dependencies from gomod
RUN go get ./...

# Build app inside container
RUN go install github.com/sk000f/metrix/cmd/web

# Run the application when the container starts
ENTRYPOINT /go/bin/web

# Container listening on port
EXPOSE 8080