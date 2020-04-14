# Use GoLang version 1.14.2-Alpine-3.11
# Intermediate compilation of project
FROM golang:1.14.2-alpine3.11 as SETUP

# Add dependencies
RUN apk add git upx

# Define working directory
WORKDIR /work/

# Add CoreDNS
RUN git clone https://github.com/coredns/coredns ./

# Download this repository as a dependency
RUN go get github.com/ddboats/big-gorilla

# Add plugin configuration file
COPY plugin.cfg ./

# Update build definitions
RUN go generate

# Build project and shrink
RUN go build -ldflags="-s -w"

# Use scratch
# Final image used for runtime
FROM scratch as RUNTIME

# Define working directory
WORKDIR /

# Copy built CoreDNS binary to image
COPY --from=SETUP /work/coredns ./

# Copy CoreDNS configuration
COPY Corefile ./

# Expose default DNS port
EXPOSE 53

# Define image entry point
ENTRYPOINT [ "/coredns" ]
