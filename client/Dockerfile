FROM node:12.16-alpine

WORKDIR /app

COPY ./package.json ./
RUN yarn
COPY . .

RUN yarn build