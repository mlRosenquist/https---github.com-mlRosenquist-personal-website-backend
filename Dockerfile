FROM node AS builder
WORKDIR /app
COPY ./package.json ./
RUN npm install --legacy-peer-deps
COPY . .
RUN npm run build


FROM node
WORKDIR /app
COPY --from=builder /app ./
CMD ["npm", "run", "start:prod"]