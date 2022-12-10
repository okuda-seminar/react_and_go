FROM --platform=arm64 node:18

COPY package-lock.json /frontend/package-lock.json
COPY package.json /frontend/package.json

WORKDIR /frontend
RUN npm install -g npm@latest
RUN npm install