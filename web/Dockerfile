FROM node:6.9.5-alpine AS build-node
WORKDIR /web
COPY . /web
RUN npm install -g -s --no-progress yarn webpack webpack-cli && \
    rm -rf /web/node_modules && \
    yarn install 
CMD webpack --hot
