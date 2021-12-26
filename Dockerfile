FROM node:12.22 AS ANGULAR_BUILD
RUN npm install -g @angular/cli@13.0.4
COPY public/ui/ /angularapp
WORKDIR angularapp
RUN npm install && ng build

FROM golang:alpine AS GO_BUILD
COPY . /
WORKDIR /
RUN ls -la
RUN go build /cmd/webapp/*.go

FROM alpine:latest
WORKDIR app
COPY --from=GO_BUILD /main ./
COPY --from=ANGULAR_BUILD /angularapp/dist/ui/* ./public/ui/dist/ui/
RUN ls -la
RUN ls -la ./public/ui/dist/ui/
CMD ./main