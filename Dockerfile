FROM golang:1.13.2-alpine AS go-build-env
RUN apk add --no-cache --update make git alpine-sdk gcc build-base

WORKDIR /src
ADD . /src

RUN go mod download
RUN make test
RUN make build

# run stage
FROM alpine 
# This App needs timezone info to work properly!
RUN apk --no-cache add tzdata
WORKDIR /src
COPY --from=go-build-env /src/bin/covid19-exporter /src
ENTRYPOINT ["./covid19-exporter"]
