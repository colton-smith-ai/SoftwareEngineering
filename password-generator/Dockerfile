###################
# Docker Commands #
###################

# Build Image: $docker build/
# -t image-name/
# .

# Run Containers: $docker run/
# --name container-name/
# --rm /
# image-name

#############################
# Building Multistage Image #
#############################

################################################
# STEP 1 Build Full Image to Compile GO Binary #
################################################

# Base image with Go pre-installed
# Alpine chosen for its small footprint
FROM golang:1.17-alpine3.15 AS buildingblock

# Create /goCode directory within
# image to hold application source files.
# Move (CD) into /goCode automatically.
# Any further commands are now executed
# inside /goCode.
WORKDIR /goCode

# Copy everything from current directory 
# into working directory (/goCode).
# Go source files from local machine are 
# now copied into image at working directory.
COPY . .

# Download Go dependencies using go.mod
RUN go mod download

# Compile binary executable of the Go app
# & save to separate WORLDIR/ouptput 
# directory with the name newPassword.
# "go build" will save binary to image root 
# (~) directory NOT WORKDIR.
# "go build" will create exe with the name 
# of module specified in go.mod, unless 
# overriden with output option (-o), 
# followed by path/to/output/exe/exe_name. 
RUN go build -o ./output/newPassword

##################################################
# STEP 2 Deploy Small Image to Execute GO Binary #
##################################################

# Start with simplist base image that just needs
# to run executable
FROM scratch

# Copy binary executable from STEP 1 image,
# with executable located at WORKDIR/output
# directory, and save executable to root 
# directory (~) of this STEP 2 image
COPY --from=buildingblock /goCode/output/newPassword /newPassword

# Entrypoint command for containers to
# run binary executable that starts app
# NOTE: CMD executes at image root (~) 
# directory NOT WORKDIR
CMD [ "/newPassword" ]