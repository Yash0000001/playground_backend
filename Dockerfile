# Use the official Alpine Linux as the base image
FROM alpine:latest

# Install necessary dependencies
RUN apk update && \
    apk add --no-cache \
    gcc \
    g++ \
    openjdk11 \
    python3 py3-pip\
    nodejs npm\
    npm \
    go \
    rust \
    cargo \
    git \
    curl \
    bash \
    build-base



# Install TypeScript globally
RUN npm install -g typescript

# Set the working directory inside the container
WORKDIR /app

# Copy your backend system files into the container
COPY . .

# Copy the .env file into the container
COPY .env .

# Expose the port your backend system will run on
EXPOSE 8080

# Command to run your backend system
CMD ["go", "run", "main.go"]