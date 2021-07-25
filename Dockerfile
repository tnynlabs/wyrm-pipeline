FROM golang:buster AS builder

RUN apt-get update -y && apt-get install -y protobuf-compiler git \
  && go get github.com/golang/protobuf/protoc-gen-go \
  && go get google.golang.org/grpc/cmd/protoc-gen-go-grpc \
  && cp /go/bin/protoc-gen-go /usr/bin/ \
  && cp /go/bin/protoc-gen-go-grpc /usr/bin/

WORKDIR /wyrm-pipeline

# Cache dependencies
# COPY go.mod .
# COPY go.sum .
# RUN go mod download

COPY . .
RUN ./scripts/build_proto.sh
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wyrm-pipeline cmd/workers_server/main.go

FROM scratch
WORKDIR /wyrm-pipeline
COPY --from=builder /wyrm-pipeline/wyrm-pipeline .
ENTRYPOINT ["./wyrm-pipeline"]
