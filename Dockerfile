FROM node:21-alpine

WORKDIR /thermal-printer

COPY package.json package.json

RUN npm i 

COPY config config
COPY index.mjs index.mjs

ENV PORT=8888
ENV PRINTER_INTERFACE=/dev/usb/lp0
ENV PRINTER_TYPE=EPSON

CMD [ "node","index.mjs" ]