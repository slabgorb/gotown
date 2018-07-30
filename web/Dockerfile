FROM node:6.9.5-alpine AS build-node
ADD ./package.json ./yarn.lock ./webpack.config.js .babelrc .eslintrc /web/
WORKDIR /web
COPY . /web
RUN npm install -g -s --no-progress yarn webpack webpack-cli && \
    cd /web && \
    yarn install --modules-folder /app/node_modules
RUN ln -s /app/node_modules /web/node_modules
RUN ln -s /web/styles /app/styles
RUN ln -s /web/index.html /app/index.html
CMD webpack --hot; tail -f /dev/null 
