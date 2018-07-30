FROM node:6.9.5-alpine AS build-node
ADD ./package.json ./yarn.lock ./webpack.config.js .babelrc .eslintrc /web/
WORKDIR /web
COPY . /web
RUN npm install -g -s --no-progress yarn webpack webpack-cli && \
    cd /web && \
    yarn install --modules-folder /app/node_modules && \
    ln -s /app/node_modules /web && \
    ln -s /web/styles /app && \
    ln -s /web/index.html /app 
CMD webpack --hot; tail -f /dev/null 
