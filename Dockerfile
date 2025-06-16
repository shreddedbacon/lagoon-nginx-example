FROM uselagoon/nginx:25.5.0

COPY index.html /app/.
COPY .lagoon.yml /app/.
