<<<<<<< HEAD
FROM debian:bookworm

WORKDIR /thermal-printer

RUN apt update 
RUN apt install libusb-1.0-0 libusb-1.0-0-dev -y
RUN apt install usbutils

COPY ./thermal-printer ./thermal-printer

ENV PORT=8888

ENV PRINTER_VENDOR=EPSON
ENV PRINTER_NAME=TM-T20X
ENV PRINTER_WIDTH=48


CMD [ "./tp" ]
=======
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
>>>>>>> main
