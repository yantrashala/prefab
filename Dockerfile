FROM golang:1.11 as gobuild

#RUN go get -d github.com/magefile/mage

#WORKDIR  $GOPATH/src/github.com/magefile/mage

#RUN go run bootstrap.go install

WORKDIR /go/src/github.com/yantrashala/prefab

COPY . ./


RUN go get -d -v
#RUN mage get

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fab .
#RUN mage prefabnogitinfo

FROM node:11-alpine as nodebuild
# Create app directory
WORKDIR /ui

COPY ./ui/package*.json ./

# Install app dependencies
RUN npm install

COPY ./ui /ui

# Build react/vue/angular bundle static files
RUN npm run build


FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=gobuild /go/src/github.com/yantrashala/prefab/fab ./fab
COPY --from=nodebuild /ui/build ./ui/build
CMD ["./fab"]