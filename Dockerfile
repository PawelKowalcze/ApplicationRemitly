# Use the official Golang image as the base image
FROM golang:1.24.0

## Install goose
#RUN apk add --no-cache git
#RUN go install github.com/pressly/goose/v3/cmd/goose@latest
#
## Set the entrypoint to run goose
#ENTRYPOINT ["goose"]

# Set the Current Working Directory inside the container
WORKDIR /usr/src

# Copy the source code into the container
COPY . .

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Build the Go app
#RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
#CMD ["./main"]