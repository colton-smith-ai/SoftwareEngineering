# Build Image: $docker build/
# -t image-name/
# .

# THIS APP REQUIES INTERACTIVE DOCKER
# CONTAINERS, (i.e. docker run -it ...)

# Run Containers: $docker run/
# --name wg-cli/
# -it/
# --rm /
# --network world-greeting-network/
# -e API_CONTAINER_NAME=wg-api/
# -e API_CONTAINER_PORT=8010/
# world-greeting-client

# Base image with Go pre-installed
# Alpine chosen for its small footprint
FROM golang:1.17-alpine3.15

# This app connects with REST API
# server. App uses ENV variables
# to identify what domain (docker
# container name) and port the
# API server is running on.
# Go files are coded to use env variable
# i.e. os.Getenv("API_CONTAINER_NAME")
# i.e. os.Getenv("API_CONTAINER_PORT")
# SPECIFY DIFFERENT ENV VAR VALUES
# WITH $docker run -e API_CONTAINER_NAME
# = container-name API_CONTAINER_PORT
# = new-port-int ...
ENV API_CONTAINER_NAME=wg-api
ENV API_CONTAINER_PORT=8010

# Create /buildClient directory within
# image to hold application source files.
# Move (CD) into /buildClient automatically.
# Any further commands are now executed
# inside WORKDIR (/buildClient).
WORKDIR /buildClient

# Copy everything in root directory into
# working directory (/buildClient).
# Go source files from local machine are
# now copied into image at working directory.
COPY . .

# Download Go dependencies using go.mod
RUN go mod download

# Compile binary executable of the Go app
# & save to separate WORKDDIR/bin directory
# with the exe named startApiClient.
# NOTE: "go build" will save binary to
# image root (~) directory NOT WORKDIR.
# NOTE: "go build" will create exe with
# the name of module specified in go.mod,
# unless overriden with output option (-o),
# followed by path/to/output/exe/exe_name.
RUN go build -o ./bin/startApiClient

# Entrypoint command for containers to
# run binary executable that starts app
# (which connects to API server)
# NOTE: CMD executes at image root (~)
# directory NOT WORKDIR
CMD [ "/buildClient/bin/startApiClient" ]