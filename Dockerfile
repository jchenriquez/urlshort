FROM golang:1.8

WORKDIR /go/src/github.com/alemjc/gophercises/urlshort

COPY . .

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/handlers
RUN go get -u gopkg.in/yaml.v2
RUN go install -v ./...

EXPOSE $PORT

CMD ["cmd"]