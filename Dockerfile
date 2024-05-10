FROM golang:1.22-bookworm AS build

RUN apt update 
RUN apt install libusb-1.0-0 libusb-1.0-0-dev -y
RUN apt install usbutils

WORKDIR /
RUN git clone https://github.com/Kruceo/thermal-printer.git

WORKDIR /build
RUN mv /thermal-printer/* .
RUN go build

FROM debian:bookworm-20240423

RUN apt update 
RUN apt install libusb-1.0-0 libusb-1.0-0-dev -y
RUN apt install usbutils

WORKDIR /thermal-printer

COPY --from=build /build/thermal-printer /thermal-printer/server

ENV PORT=8888

ENV PRINTER_VENDOR=EPSON
ENV PRINTER_NAME=TM-T20X
ENV PRINTER_WIDTH=48
ENV PRINTER_CHARSET=WPC1254_Turkish

CMD [ "./server" ]