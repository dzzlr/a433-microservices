FROM node:14.21-alpine as base
WORKDIR /app
COPY package*.json ./
 
FROM base as dev
RUN apk add --no-cache bash
RUN wget -O /bin/wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh
RUN chmod +x /bin/wait-for-it.sh
 
ENV NODE_ENV=development
ENV PORT=3001
RUN npm install
COPY ./*.js ./
EXPOSE 3001
CMD ["npm", "start"]