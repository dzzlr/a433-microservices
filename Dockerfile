FROM node:18-alpine
WORKDIR /src
COPY . .
RUN npm install

ENV PORT=3001
EXPOSE 3001

CMD ["npm","start"]