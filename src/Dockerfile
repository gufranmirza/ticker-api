FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GOOS=linux \
    GOARCH=amd64\
    BINARY_NAME="ticker-api"

# Move to working directory /build
WORKDIR $GOPATH/src/github.com/gufranmirza/ticker-api/src

# Copy the code into the container
COPY . .

# Download dependency using go mod
RUN go mod download

# Build Project 
RUN env GOOS=linux GOARCH=amd64 go build -o ../bin/${BINARY_NAME} .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp $GOPATH/src/github.com/gufranmirza/ticker-api/bin/${BINARY_NAME} .
RUN mv $GOPATH/src/github.com/gufranmirza/ticker-api/src/app-config-docker.json app-config.json 

# Export necessary port
EXPOSE 8001

# Command to run when starting the container
CMD ["/dist/ticker-api", "serve"]
