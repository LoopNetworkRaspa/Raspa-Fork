# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/LoopNetworkRaspa/Raspa-Fork

WORKDIR /go/src/github.com/LoopNetworkRaspa/Raspa-Fork

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache kaspad dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o kaspad .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/LoopNetworkRaspa/Raspa-Fork/kaspad /app/
COPY --from=build /go/src/github.com/LoopNetworkRaspa/Raspa-Fork/infrastructure/config/sample-kaspad.conf /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
