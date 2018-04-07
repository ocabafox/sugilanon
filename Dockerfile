FROM golang
LABEL maintainer="Alexjander T. Bacalso <info.sugilanon@gmail.com>"

WORKDIR /go/src/github.com/XanderDwyl/sugilanon/

COPY assets assets
COPY app app
COPY glide.yaml .
COPY main.go .

RUN curl https://glide.sh/get | sh
RUN glide install

# building binary go
RUN CGO_ENABLED=0 go build -o webapp main.go

FROM scratch
COPY --from=0 /go/src/github.com/XanderDwyl/sugilanon/ .

EXPOSE 80 3000

ENTRYPOINT ["/webapp"]
