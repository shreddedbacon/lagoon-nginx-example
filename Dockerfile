FROM uselagoon/nginx:25.5.0
#comment
COPY index.html /app/.
COPY .lagoon.yml /app/.
