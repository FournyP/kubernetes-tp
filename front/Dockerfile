FROM node:16.17.1

WORKDIR /app

COPY ./package.json ./yarn.lock ./

RUN npm i

COPY . ./

EXPOSE 5173

CMD [ "npm", "run", "dev" ]