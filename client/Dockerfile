FROM node:lts-alpine

WORKDIR /client
COPY package.json ./
COPY yarn.lock ./
RUN yarn
COPY . .

EXPOSE 8080
EXPOSE 8000

CMD [ "yarn", "serve" ]