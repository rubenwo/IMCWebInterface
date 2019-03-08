FROM node:8
ENV NODE_ENV=production

# Create app directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# Install app dependencies
COPY ./service.api-gateway/package.json /usr/src/app/
RUN npm install

# Bundle app source
COPY ./service.api-gateway /usr/src/app

CMD [ "npm", "run", "server" ]