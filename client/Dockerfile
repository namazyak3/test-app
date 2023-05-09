FROM node:20.1.0-alpine3.17
WORKDIR /client
COPY . .
RUN npm i
CMD ["npm", "run", "dev"]
EXPOSE 3000