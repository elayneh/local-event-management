FROM node:18 AS build
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install
COPY . .
RUN yarn build

# Stage 2: Serve the Nuxt app with Nginx
FROM nginx:alpine
COPY --from=build /app/.output/public /usr/share/nginx/html

EXPOSE 3000
CMD ["nginx", "-g", "daemon off;"]
