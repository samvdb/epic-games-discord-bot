FROM golang:1.14-alpine

RUN apk --update upgrade
RUN apk add --update gcc musl-dev


# removing apk cache
RUN rm -rf /var/cache/apk/*

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/samvdb/epic-games-discord-bot
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package and create test binary
RUN mkdir -p /build
RUN GOARCH=amd64 CGO_ENABLED=1 go build  -a -ldflags "-linkmode external -extldflags -static"    -o /build/bot .
# Perform any further action as an unprivileged user.
USER nobody:nobody



# Run the executable
CMD ["/build/bot"]