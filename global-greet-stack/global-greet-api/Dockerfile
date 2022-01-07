# Build Image: $docker build/
# -t image-name/
# .

# Run Containers: $docker run/
# --name globalgreet-api/
# -e API_PORT=8010/
# -e DB_CONTAINER_NAME=globalgreet-db/
# -e DB_PORT=6379/
# -p 8010:8010/
# --rm/
# --network global-greet-network/
# global-greet-api

# Base image with Go pre-installed
# Alpine chosen for its small footprint
FROM golang:1.17-alpine3.15

# Set environment varialbes for:
#   - Server Port for API
#   - Database Container Name
#   - Database Container Port
# Go files are coded to use env variable
# i.e. os.Getenv("API_PORT")
ENV API_PORT=8010
ENV DB_CONTAINER_NAME=globalgreet-db
ENV DB_PORT=6379

# Create /buildApi directory within
# image to hold application source files.
# Move (CD) into /buildApi automatically.
# Any further commands are now executed
# inside WORKDIR (/buildApi).
WORKDIR /buildApi

# Copy everything in root directory into
# working directory (/buildApi).
# Go source files from local machine are
# now copied into image at working directory.
COPY . .

# Download Go dependencies using go.mod
RUN go mod download

# Compile binary executable of the Go app
# & save to separate WORKDDIR/bin directory
# with the name startApiServer.
# NOTE: "go build" will save binary to image
# root (~) directory NOT WORKDIR.
# NOTE: "go build" will create exe with
# the name of module specified in go.mod,
# unless overriden with output option (-o),
# followed by path/to/output/exe/exe_name.
RUN go build -o ./bin/startApiServer

# Tell developer & docker what port server
# runs on
EXPOSE ${API_PORT}

# Entrypoint command for containers to
# run binary executable that starts app
# (which is RESTful API Server)
# NOTE: CMD executes at image root (~)
# directory NOT WORKDIR
CMD [ "/buildApi/bin/startApiServer" ]