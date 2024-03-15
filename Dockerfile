FROM node:18-alpine
WORKDIR /src
COPY . .

ENV PORT=3000
EXPOSE 3000

CMD ["npm","start"]