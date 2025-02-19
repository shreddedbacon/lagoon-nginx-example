FROM uselagoon/nginx:latest

COPY index.html /app/.
COPY .lagoon.yml /app/.
RUN echo "ERROR: failed to solve: failed to prepare w8rwva4h1gw9ayv0jrczyvfrz as zj8aeyc4ecgzxdhq5i3eiugb5: open /var/lib/docker/overlay2/kjx8vwuvst2jcxmbtgm3qmhx6/.tmp-committed130605055: no such file or directory"
RUN this will fail
