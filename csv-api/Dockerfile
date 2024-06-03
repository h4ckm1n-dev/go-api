# Start a new stage from scratch
FROM alpine:3.20.0

# Create a non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the Current Working Directory
WORKDIR /home/appuser/

# Copy the pre-built binary file to the container
COPY main ./
COPY data ./data

# Change ownership of the files to the non-root user
RUN chown -R appuser:appgroup /home/appuser/

# Switch to non-root user
USER appuser

# Expose port 8088 to the outside world
EXPOSE 8088

# Command to run the executable
CMD ["./main"]

