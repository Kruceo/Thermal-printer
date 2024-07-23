# Thermal Printer Server

This works only with USB printers, (tested only with EPSON TM-T20X).

## Requirements 

- libusb-1.0-0
- libusb-1.0-0-dev

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

#### Example:

```js
const res = await fetch("http://localhost:8888/get/width");
const data = await res.json();
console.log(data)
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

#### Example:

```js
const body = {
    query:[
        ["println","Hello World"],
        ["println","This is a simple test!"],
        ["println",""],
        ["cut"]
    ]
}

const options = {
    method:"POST",
    headers:{"Content-Type":"application/json"},
    body:JSON.stringify(body)
} 

const res = await fetch("http://localhost:8888/query",options);
console.log(res.status)
```

