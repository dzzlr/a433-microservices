FROM node:18-alpine
WORKDIR /src
COPY package*.json ./
RUN npm ci
COPY ./*.js ./

ENV PORT=3001
EXPOSE 3001

CMD ["npm","start"]