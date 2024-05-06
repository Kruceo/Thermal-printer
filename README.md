<<<<<<< HEAD
# Thermal Printer Server

This works only with USB printers, (tested only with EPSON TM-T20X).

## Requirements 
=======
### Thermal Printer Web API

#### Description

The Thermal Printer Web API provides an endpoint for printing text on a thermal printer. The thermal printer needs to be connected to the network or the device where the service is running.
>>>>>>> main

- libusb-1.0-0
- libusb-1.0-0-dev

<<<<<<< HEAD
## Configuration

Use `.env` file to configure the server.

```env
PRINTER_VENDOR=EPSON
PRINTER_NAME=TM-T20X
PRINTER_WIDTH=48
PRINTER_CHARSET=WPC1254_Turkish
```

## Tips

### List Devices

You can use `usbutils` to easily reach your device vendor and name.
Using the command `lsusb`.

### Running in Windows WSL

You will like to use `usbip` in windows WSL to run this in a Windows Platform

## Containers 

This works in containers, you will like to pass your USB hub with `--device` option in docker.

```bash
docker run --device=/dev/bus/usb:/dev/bus/usb rafola/thermal-printer
```

Extended:

```bash
docker run -p 8888:8888 --device=/dev/bus/usb:/dev/bus/usb -e PRINTER_CHARSET=WPC1254_Turkish -e PRINTER_WIDTH=48 -e PRINTER_VENDOR=EPSON -e PRINTER_NAME=TM-T20X rafola/thermal-printer
```

## Endpoints 

- GET /get/width 
- POST /query

### ***get/width***

Returns the printer previous configured printer width.

```json
{ "width":48 }
```

### ***/query***

Prints the requested form.

#### POST FORM:

```json
{
    "query":[
        ["center"],
        ["println","Hello world"],
        ["left"],
        ["println","Ass: Kruceo"],
        ["cut"]
    ]
}
=======
- **URL**: `POST /query`
- **Description**: This endpoint receives a list of commands and texts to be printed and sends them to the thermal printer. The commands are provided in the request body as a list of strings.

#### Request Body (Example)

```json
{
"query": [
    [ "center" ],
    [ "println", "Transaction 47" ],
    [ "left" ],
    [ "println", "------------------------------------------------" ],
    [ "println", "Type: Entry" ],
    [ "println", "Boat: SSC Tuatara" ],
    [ "println", "Supplier: Pennsylvania" ],
    [ "code128", "https://kruceo.com"],
    [ "qrcode" , "https://kruceo.com"],
    [ "cut" ]
  ]
}
```

#### Docker Command for Printer Connection

To run the printer connection in a Docker container, use the following command:

##### Linux

```bash
docker run --rm -it --device=/dev/usb/lp0 rafola/thermal-printer
```

##### Windows

```bash
docker run --rm -it --device=//./COM0 rafola/thermal-printer
>>>>>>> main
```