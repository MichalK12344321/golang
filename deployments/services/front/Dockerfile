FROM node:20.14.0-buster-slim AS build
COPY cmd/front/ .
RUN npm install
ENV VITE_SCHEDULER_URL=/api/scheduler
ENV VITE_CDG_URL=/api/cdg
ENV VITE_WEB_URL=/api/web
RUN npm run build

FROM nginx:stable-alpine3.19-slim AS final
COPY --from=build /dist /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf
COPY deployments/services/front/nginx.conf /etc/nginx/conf.d
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]