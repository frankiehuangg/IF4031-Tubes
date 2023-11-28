FROM node:21

WORKDIR /app

COPY package*.json ./

RUN npm ci

COPY . .

EXPOSE 8000

RUN rm -rf node_modules

CMD [ "npm", "run", "dev" ]
