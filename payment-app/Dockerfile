FROM node:16-alpine

WORKDIR /app

COPY package*.json ./

RUN rm -rf node_modules
RUN npm install
COPY . .
EXPOSE 8000

CMD [ "npm", "run", "dev" ]
