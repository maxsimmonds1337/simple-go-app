FROM arm64v8/golang:latest

# Install additional dependencies
RUN apt-get update && apt-get install -y gcc git wget tar xz-utils make libc6-dev

# Install IBM Data Server Driver Package dependencies
RUN apt-get install -y libxml2 libssl1.1 libaio1

# Update PATH environment variable
ENV PATH="/usr/local/go/bin:${PATH}"

# Install go_ibm_db installer
RUN go install github.com/ibmdb/go_ibm_db/installer@v0.4.2
# or use the latest version:
# RUN go install github.com/ibmdb/go_ibm_db/installer@latest

# Set environment variables
ENV DB_HOSTNAME=localhost
ENV DB_PORT=50500
ENV DB_DATABASE=hotdogs
ENV DB_USERNAME=db2inst1
ENV DB_PASSWORD=hotdog
ENV CGO_CFLAGS=-I/Users/max/go/pkg/mod/github.com/ibmdb/clidriver/include
ENV CGO_LDFLAGS=-L/Users/max/go/pkg/mod/github.com/ibmdb/clidriver/lib

# Set working directory for your application
WORKDIR /go/src/github.com/maxsimmonds1337/simple-go-app/

# Copy application code
COPY . .

# Build and run your application
RUN go build -o app ./vote/cmd

CMD ["./app"]
