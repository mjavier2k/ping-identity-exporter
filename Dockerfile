FROM golang:alpine3.12 as build

WORKDIR /src

COPY . . 

RUN apk add --update make git && \
        make build_static

FROM golang:alpine3.12

WORKDIR /usr/src/app

COPY --from=build /src/bin/ . 

COPY config.yaml.example config.yaml

EXPOSE 9999

ENTRYPOINT [ "/usr/src/app/ping-identity-exporter" ]