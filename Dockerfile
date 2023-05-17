FROM golang:latest

# Set environment variables
ENV DB_HOSTNAME=localhost
ENV DB_PORT=50500
ENV DB_DATABASE=hotdogs
ENV DB_USERNAME=db2inst1
ENV DB_PASSWORD=hotdog

ENV IBM_DB_HOME=/go/pkg/mod/github.com/ibmdb/clidriver

ENV CGO_CFLAGS=-I/go/pkg/mod/github.com/ibmdb/clidriver/include
ENV CGO_LDFLAGS=-L/go/pkg/mod/github.com/ibmdb/clidriver/lib
ENV LD_LIBRARY_PATH=$IBM_DB_HOME/lib:$LD_LIBRARY_PATH
ENV PATH=$PATH:$IBM_DB_HOME/bin

# Install additional dependencies
RUN apt-get update && apt-get install -y wget tar libxml2

# Install go_ibm_db installer
RUN go install github.com/ibmdb/go_ibm_db/installer@latest

# Change working directory to the installer location
WORKDIR /go/pkg/mod/github.com/ibmdb/go_ibm_db@latest/installer

# Download and extract the odbc_cli package
RUN wget https://public.dhe.ibm.com/ibmdl/export/pub/software/data/db2/drivers/odbc_cli/linuxx64_odbc_cli.tar.gz && \
    tar -xzf linuxx64_odbc_cli.tar.gz -C ../../

# Set working directory for your application
WORKDIR /go/src/github.com/maxsimmonds1337/simple-go-app/

# Copy application code
COPY . .

# Build your application
RUN go build -o app ./vote/cmd

# Run your application
CMD ["./app"]
