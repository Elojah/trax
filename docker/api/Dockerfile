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
WORKDIR /api
COPY . .
RUN make tidy
RUN env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 make api

FROM scratch AS runner

WORKDIR /app
COPY --from=builder --chmod=755 /api/bin/trax_api .
COPY --from=builder /api/config/api/docker.json .
COPY --from=builder /api/config/api/private.key .
COPY --from=builder /api/config/api/public.crt .
ENTRYPOINT ["/app/trax_api", "/app/docker.json"]
