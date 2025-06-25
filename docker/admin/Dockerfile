FROM alpine:latest AS builder

# Install Go and make
RUN apk add --no-cache \
	go \
	make

# Set Go environment variables
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
ENV GO111MODULE=on

# Create workspace directory
WORKDIR /admin
COPY . .
RUN make tidy
RUN env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 make admin

FROM scratch AS runner

WORKDIR /app
COPY --from=builder --chmod=755 /admin/bin/trax_admin .
COPY --from=builder /admin/config/admin/local.json .
ENTRYPOINT ["/app/trax_admin", "/app/local.json"]
