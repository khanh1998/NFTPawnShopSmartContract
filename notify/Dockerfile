FROM node:14-alpine

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 7789

CMD ["npm", "run", "start"]
