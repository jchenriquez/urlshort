FROM node:12 AS UI

WORKDIR /go/src/github.com/alemjc/gophercises/urlshort/web
RUN npm install yarn -g

COPY ./web/urlshort .

RUN yarn build

FROM golang:1.8

WORKDIR /go/src/github.com/alemjc/gophercises/urlshort

COPY ./cmd ./cmd
COPY ./rest ./rest
COPY --from=UI /go/src/github.com/alemjc/gophercises/urlshort/web ./web

RUN ls ./*

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/handlers
RUN go get -u github.com/boltdb/bolt/...
RUN go get -u gopkg.in/yaml.v2
RUN go install -v ./cmd
RUN go install -v ./rest

EXPOSE $PORT

CMD ["cmd"]

