# Use an official Golang runtime as the base image
FROM golang:alpine AS build

# Define a build argument for VERSION
ARG VERSION

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download and install project dependencies
RUN go mod download

# Copy the entire project to the container
COPY . .

# install templ
RUN go install github.com/a-h/templ/cmd/templ@latest

# Generate templates
RUN templ generate

# Build the Go application
RUN GOOS=linux go build -o app .

# Create a minimal runtime image
FROM alpine

# Copy the binary from the build stage to the runtime image
COPY --from=build /app/app /app
COPY --from=build /app/view /view


# Expose the port that your Go application listens on
EXPOSE 8080

CMD ["./app"]
