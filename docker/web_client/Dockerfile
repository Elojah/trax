FROM alpine/git:latest AS builder

# Install Go and make
RUN apk add --no-cache \
	go \
	make

# Install npx
RUN apk add --no-cache \
	nodejs \
	npm

# Set Go environment variables
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
ENV GO111MODULE=on

# Create workspace directory
WORKDIR /web_client
COPY . .
# Build the static client
ENV CONFIG_FILE=/web_client/config/client/docker.json
RUN make npm
RUN make client
# Build the web client
RUN make tidy
RUN env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 make web_client

FROM scratch AS runner

WORKDIR /app
COPY --from=builder --chmod=755 /web_client/bin/client/static .
COPY --from=builder --chmod=755 /web_client/bin/trax_web_client .
COPY --from=builder /web_client/config/web_client/docker.json .
COPY --from=builder /web_client/config/web_client/private.key .
COPY --from=builder /web_client/config/web_client/public.crt .
ENTRYPOINT ["/app/trax_web_client", "/app/docker.json"]
