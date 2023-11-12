FROM golang:1.21.4-alpine as builder

WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=$(go env GOOS) GOARCH=$(go env GOARCH) go build -o /exe

FROM scratch
COPY --from=builder /exe /e-boerse
ENTRYPOINT [ "/e-boerse" ]
