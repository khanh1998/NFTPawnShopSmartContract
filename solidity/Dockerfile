FROM node:14-alpine

WORKDIR /usr/src/app

COPY package.json ./

RUN npm install

COPY . .

CMD ["npx","truffle", "migrate", "--network", "ganachecli"]