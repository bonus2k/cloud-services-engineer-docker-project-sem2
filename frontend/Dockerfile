FROM node:16 AS builder
WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .
RUN VUE_APP_API_URL=/api/ npm run build

FROM nginx:1.25-alpine-slim

ARG APP_USER=appuser
ARG APP_UID=1000
ARG APP_GID=1000
RUN addgroup -g ${APP_GID} ${APP_USER}  && adduser -u ${APP_UID} -G ${APP_USER} -s /bin/sh -D ${APP_USER}
RUN mkdir -p /var/cache/nginx/client_temp && chown -R ${APP_UID}:${APP_GID} /var/cache/nginx

USER ${APP_UID}:${APP_GID}

COPY ./nginx.conf /etc/nginx/nginx.conf
COPY --chown=${APP_UID}:${APP_GID} --from=builder /app/dist /usr/share/nginx/html/momo-store/

EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]