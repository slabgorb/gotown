FROM node:6.9.5-alpine AS build-node
ADD ./package.json ./yarn.lock ./webpack.config.js .babelrc .eslintrc /web/
RUN npm install -g -s --no-progress yarn && \
    cd /web && \
    yarn 

RUN npm install webpack -g -s
RUN npm install webpack-cli -g -s

# RUN mkdir -p /opt/app && cp -a /tmp/node_modules /opt/app/
WORKDIR /web
COPY . /web

CMD webpack --hot  --mode=development