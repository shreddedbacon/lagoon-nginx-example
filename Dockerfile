FROM uselagoon/nginx:25.5.0

COPY index.html /app/.
COPY allowed.html /app/.
COPY restricted.html /app/.
COPY .lagoon.yml /app/.
COPY app.conf /etc/nginx/conf.d/app.conf