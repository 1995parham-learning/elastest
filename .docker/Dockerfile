FROM golang:1.12.1 as builder

# Modify these variables depending on your application:
# * `TZ` Timezone for the app.
# * `GO_SOURCE_PATH` The path to your Go source code.
# * `GO_BINARY` The main app binary that is built.
# * `GOROOT` This should not need changed.
# * `GOPATH` This should not n  eed changed.
# * `GOBIN` This should not need changed.
ENV TZ=Asia/Tehran \
    GO_SOURCE_PATH=/app/src/gitlab.snapp.ir/golangify/elastest \
    GO_BINARY="elastest" \
    GO111MODULE=on

# Handles goproxy in build arguments
ARG GOPROXY
ENV GOPROXY ${GOPROXY}

# Install the base required packages for the image. Find any additional packages
# from [the Alpine package explorer](https://pkgs.alpinelinux.org/packages).
RUN apt-get update && \
    apt-get install -y \
      git \
      ca-certificates

# Set the current directory for the Docker image to the Go source path, since
# we initially want to build the Go app.
WORKDIR $GO_SOURCE_PATH

# Copy the application files.
COPY . .

# Install the packages required for building the Go application.
RUN go mod download

# Build the application binary, if you have additional binaries add them below.
RUN CGO_ENABLED=0 GOOS=linux go build -o $GO_BINARY .

FROM alpine:3.9

ENV TZ=Asia/Tehran \
    GO_SOURCE_PATH=/app/src/gitlab.snapp.ir/golangify/elastest \
    GO_BINARY="elastest" \
    GOPATH=/app \
    GOBIN=/app/bin

WORKDIR $GOBIN

COPY --from=builder $GO_SOURCE_PATH .

RUN apk add --update --no-cache \
      tzdata \
      ca-certificates \
      bash \
    && \
    cp --remove-destination /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo "${TZ}" > /etc/timezone && \
    mkdir -p /var/log && \
    chgrp -R 0 /var/log && \
    chmod -R g=u /var/log

# Specify the command to run when the container starts.
CMD ["${GOBIN}/${GO_BINARY}"]

# Expose any port your application listens on. You must use ports higher than 1024.
EXPOSE 8080
