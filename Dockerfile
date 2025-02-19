FROM uselagoon/nginx:latest

COPY index.html /app/.
COPY .lagoon.yml /app/.
RUN echo "ERROR: failed to solve: layer does not exist"
RUN this will fail
