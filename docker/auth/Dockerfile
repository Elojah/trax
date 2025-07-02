FROM alpine/git:latest AS builder

# Install Go and make
RUN apk add --no-cache \
	go \
	make

# Set Go environment variables
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
ENV GO111MODULE=on

# Create workspace directory
WORKDIR /auth
COPY . .
RUN make tidy
RUN env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 make auth

FROM scratch AS runner

WORKDIR /app
COPY --from=builder --chmod=755 /auth/bin/trax_auth .
COPY --from=builder /auth/config/auth/docker.json .
ENTRYPOINT ["/app/trax_auth", "/app/docker.json"]
