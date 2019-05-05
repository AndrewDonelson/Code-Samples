FROM golang:latest as builder

# install xz
RUN apt-get update && apt-get install -y \
    xz-utils \
    && rm -rf /var/lib/apt/lists/*

# install UPX
ADD https://github.com/upx/upx/releases/download/v3.95/upx-3.95-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.95-amd64_linux.tar.xz | \
    tar -xOf - upx-3.95-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx

# Create and change into our target source code folder
RUN mkdir --parents /go/src/github.com/NlaakStudiosLLC/Company-Website/client \
    && cd /go/src/github.com/NlaakStudiosLLC/Company-Website/client

# Set the default working folder and add our source
# add everything not ignored by .dockerignore
WORKDIR /go/src/github.com/NlaakStudiosLLC/Company-Website/client
ADD . ./

# Build it
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

# Strip and Compress the binary
RUN strip --strip-unneeded main
RUN upx main

# Build image to be deployed
FROM scratch
COPY --from=builder /go/src/github.com/NlaakStudiosLLC/Company-Website/client/main /go/src/app/main
COPY --from=builder /go/src/github.com/NlaakStudiosLLC/Company-Website/client/static /go/src/app/static
WORKDIR /go/src/app
ENTRYPOINT [ "./main" ]
