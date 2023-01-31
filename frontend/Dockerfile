FROM --platform=arm64 node:16.17.0

COPY package-lock.json /frontend/package-lock.json
COPY package.json /frontend/package.json

WORKDIR /frontend
RUN npm install 
