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
ENV PRINTER_CHARSET=WPC1254_Turkish

CMD [ "./thermal-printer" ]