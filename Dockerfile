FROM uselagoon/nginx:latest

COPY index.html /app/.
COPY .lagoon.yml /app/.

RUN this should fail
