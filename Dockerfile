FROM golang
LABEL maintainer="Alexjander T. Bacalso <info.sugilanon@gmail.com>"

WORKDIR /go/src/github.com/XanderDwyl/sugilanon/
COPY vendor vendor
COPY assets assets
COPY app app
COPY main.go .
RUN ls -al
RUN CGO_ENABLED=0 go build -o webapp main.go

FROM scratch
COPY --from=0 /go/src/github.com/XanderDwyl/sugilanon/ .

EXPOSE 3000

ENTRYPOINT ["/webapp"]
