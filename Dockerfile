FROM golang:alpine as builder
ENV GO111MODULE=on
WORKDIR /build
ADD ./ /build
RUN apk update --no-cache
RUN apk add git
RUN go build -o eks_bp_demo  .
# Can run smoke unit test for build verification

FROM golang:alpine
WORKDIR /app
COPY --from=builder /build/eks_bp_demo .
EXPOSE 8080
ENTRYPOINT ["/app/eks_bp_demo"]