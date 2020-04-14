# Use GoLang version 1.14.2-Alpine-3.11
# Intermediate compilation of project
FROM golang:1.14.2-alpine3.11 as SETUP

# Add dependencies
RUN apk add git upx

# Define working directory
WORKDIR /work/

# Add CoreDNS
RUN git clone https://github.com/coredns/coredns ./

# Copy go project
COPY ./src /project

# Tell go where to find the project when CoreDNS imports it
RUN go mod edit -replace github.com/ddboats/big-gorilla=/project/

# Add plugin configuration file
COPY plugin.cfg ./

# Update build definitions
RUN go generate

# Build project and shrink
RUN go build -ldflags="-s -w"

# Use scratch
# Final image used for runtime
FROM alpine:3.11.5 as RUNTIME

WORKDIR /

# Copy built CoreDNS binary to image
COPY --from=SETUP /work/coredns ./

# Copy CoreDNS configuration
COPY Corefile ./

# Expose default DNS port
EXPOSE 53

# Define image entry point
ENTRYPOINT [ "/coredns" ]
