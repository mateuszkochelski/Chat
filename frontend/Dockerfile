FROM node:20-alpine

WORKDIR /app

COPY frontend/package*.json ./

RUN npm install

COPY . .

RUN cd frontend && npm run build

EXPOSE 4173

CMD ["npm", "run", "preview"]
