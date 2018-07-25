FROM node:6.9.5-alpine AS build-node
ADD ./package.json ./yarn.lock ./webpack.config.js /opt/app/
RUN npm install -g -s --no-progress yarn && \
    mkdir -p /opt/app && \
    cd /opt/app && \
    yarn 

RUN npm install webpack -g -s
RUN npm install webpack-cli -g -s

# RUN mkdir -p /opt/app && cp -a /tmp/node_modules /opt/app/
WORKDIR /opt/app
COPY . /opt/app

CMD webpack --hot