FROM node:15.14.0

ENV ROOT=/src/client
WORKDIR ${ROOT}

RUN apt update && \
 apt add git && \
 npm install -g npm && \
 npm install -g yarn

COPY ./client/package.json ./
COPY ./client/yarn.lock ./
RUN yarn install

COPY ./client .
RUN yarn run build

EXPOSE 3000