FROM node:16 AS builder
WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .
RUN VUE_APP_API_URL=/api/ npm run build

FROM nginx:1.25-alpine-slim
COPY --from=builder /app/dist /usr/share/nginx/html/momo-store/
COPY ./nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]