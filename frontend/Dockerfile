FROM node:16-slim as BUILDER


WORKDIR build
COPY . .

RUN npm install
RUN npm run build -- --prod

FROM nginx:1.21-alpine

COPY --from=BUILDER /build/dist/frontend/ /usr/share/nginx/html/
