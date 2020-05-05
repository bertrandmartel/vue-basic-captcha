# build environment
FROM node:12.2.0-alpine as frontend
WORKDIR /app
COPY frontend/package.json /app/package.json
RUN npm install
COPY frontend/vue.config.js /app/vue.config.js
COPY frontend/src /app/src
COPY frontend/public /app/public
RUN npm run build


FROM golang:alpine
RUN mkdir /app 
ADD . /app
WORKDIR /app/backend
COPY --from=frontend /app/dist /app/backend/dist
RUN go build -o main .
WORKDIR /app/provisioning/add
RUN go get ./...
RUN go build -o add .
WORKDIR /app/provisioning/delete
RUN go get ./...
RUN go build -o delete .
WORKDIR /app/backend
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./main"]
