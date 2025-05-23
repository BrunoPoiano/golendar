FROM golang:1.24.2-alpine3.21 AS builder

# Set the working directory for the build stage
WORKDIR /build

# Copy only files needed for building
COPY . .

# Build the Go application with static linking to reduce dependencies
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o golendar .

# Use a smaller Alpine base for the final image
FROM alpine:3.21

# Set the working directory for the application
WORKDIR /app/golendar

# Install only the necessary packages in a single RUN to reduce layers
RUN apk add --no-cache dcron tzdata && \
  echo "0 8 * * * /app/golendar/golendar >> /var/log/cron.log 2>&1" > /etc/crontabs/root && \
  touch /var/log/cron.log

# Copy only the compiled binary from the builder stage
COPY --from=builder /build/golendar ./

# Copy startup script and make it executable
COPY start.sh ./
RUN chmod +x ./start.sh

# Execute the startup script when container launches
CMD ["/bin/sh", "./start.sh"]
