FROM node:15.14.0
RUN apt update && apt upgrade -y

ENV ROOT=/src/client
WORKDIR ${ROOT}

COPY ./client/package.json ./
COPY ./client/yarn.lock ./
RUN yarn install

COPY ./client .
RUN yarn build

CMD ["yarn", "start"]