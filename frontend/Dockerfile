FROM node:18-alpine AS builder

# Add a work directory
WORKDIR /app

COPY package.json .
COPY yarn.lock .
COPY entrypoint.sh .

RUN yarn install --network-timeout 1000000

COPY . .

RUN chmod +x entrypoint.sh

RUN yarn build
RUN chgrp -R 0 /app && \
    chmod -R g=u /app

EXPOSE 8080

ENTRYPOINT [ "/app/entrypoint.sh" ]
CMD [ "sh" ]
