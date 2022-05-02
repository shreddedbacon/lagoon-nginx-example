FROM amazeeio/nginx

COPY index.html /app/.
COPY .lagoon.yml /app/.