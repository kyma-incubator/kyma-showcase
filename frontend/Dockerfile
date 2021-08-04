FROM alpine:3.14.0 AS builder
# install node
RUN apk add --update nodejs npm

# set working directory
WORKDIR /app

# install app dependencies
COPY ./package.json /app
COPY ./package-lock.json /app
RUN npm ci
COPY ./src /app/src 
COPY ./public /app/public 
COPY ./jsconfig.json /app 
RUN npm run build

## Main stage ##
FROM nginx:1.12.1
COPY ./nginx.conf /etc/nginx/nginx.conf
COPY --from=builder /app/build /etc/nginx/app

EXPOSE 80