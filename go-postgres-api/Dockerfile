
FROM alpine:3.20.0

# Create the app user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Install glibc
RUN apk --no-cache add curl && \
  curl -Lo /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub && \
  curl -Lo glibc-2.34-r0.apk https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.34-r0/glibc-2.34-r0.apk && \
  apk add --force-overwrite glibc-2.34-r0.apk && \
  rm glibc-2.34-r0.apk

# Set the working directory
WORKDIR /home/appuser/

# Copy the application files
COPY go-postgres-api ./

# Ensure the script has execute permissions
RUN chown -R appuser:appgroup /home/appuser/ && chmod +x ./csv-insert

# Switch to the non-root user
USER appuser

# Expose the application port
EXPOSE 8088

# Set the entrypoint to the csv-insert script
ENTRYPOINT ["./go-postgres-api"]

