# build stage
FROM golang:1.8.1-alpine AS build-env
ADD . /src
RUN cd /src && go build -a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo -o goapp .

# final stage
FROM scratch
WORKDIR /app
COPY --from=build-env /src/goapp /app/
EXPOSE 8080
ENTRYPOINT ["./goapp"]
