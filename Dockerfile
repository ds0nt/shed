# Use a base image with Go installed
FROM golang:latest AS builder

# Set the working directory
WORKDIR /app

# Copy the main.go file to the working directory
COPY . .

# Build the Go application
RUN go build -o app .

# Use a separate stage for building the UI
FROM node:latest AS ui-builder

# Set the working directory for the UI
WORKDIR /app/ui

# Copy the UI source code to the working directory
COPY ui/shed/dist .

# Build the UI
RUN npm install

# Use a lightweight base image for the final image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built Go application from the builder stage
COPY --from=builder /app/app .

# Copy the built UI from the ui-builder stage
COPY --from=ui-builder /app/ui .

# Expose any necessary ports
EXPOSE 8080

# Set the command to run the application
CMD ["./app"]