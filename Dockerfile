# Use the official Golang image as the base image
FROM golang:1.24.0

# Set the Current Working Directory inside the container
WORKDIR /usr/src

# Copy the source code into the container
COPY . .

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy


# Expose port 8080 to the outside world
EXPOSE 8080

