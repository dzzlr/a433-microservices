FROM node:18-alpine
WORKDIR /src
COPY package*.json ./
RUN npm ci
COPY ./*.js ./

ENV PORT=3000
EXPOSE 3000

CMD ["npm","start"]