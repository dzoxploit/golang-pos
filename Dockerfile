# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/go-pos

# Copy the entire project to the working directory
COPY . .

# Download the dependencies
RUN go get -d -v ./...

# Install the application
RUN go install -v ./...

# Set the default command to run your binary
CMD ["go-pos"]
