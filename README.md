# Thermal Printer Server

This works only with USB printers, (tested only with EPSON TM-T20X).

## Requirements 

- **Linux**
    - golang
    - libusb-1.0-0
    - libusb-1.0-0-dev
- **Windows**
    - golang
    - wsys2 (optional)
    - libusb-1.0
    - gcc
    - pkg-config
    - zadig (optional)
## Running From Source on Linux

On Linux, setting up to run this program is straightforward. Just install `libusb-1.0` on your distribution, use `go build` to build the program, and then execute the binary.

## Running From Source on Windows

To run this program on Windows, you’ll need to install `gcc`, `pkg-config`, and `libusb-1.0`. It’s recommended to use **MinGW** or a similar tool to install these dependencies. In this example, we’ll use [MSYS2](https://www.msys2.org/) for easy setup, although there are other methods available.

> This guide assumes MSYS2 is installed at `C:/msys2`.

**In an MSYS2 shell:**

```shell
pacman -S mingw-w64-ucrt-x86_64-gcc
pacman -S mingw-w64-x86_64-pkg-config
pacman -S mingw-w64-x86_64-libusb
```

### Important Steps

After installing these packages, you’ll need to update your Windows environment variables:

1. Add `C:/msys2/ucrt/bin` and `C:/msys2/mingw64/bin` to the `PATH` environment variable.
2. If `PKG_CONFIG_PATH` doesn’t exist, create it and set it to `C:/msys2/mingw64/lib/pkgconfig`.

This setup allows you to successfully run `go build`.

### Troubleshooting Device Detection

If the program doesn’t recognize your device, it may be due to driver issues. You can use the [Zadig](https://zadig.akeo.ie/) tool to replace the device driver with *WinUSB* or *libusbK*, then try running the program again.

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

### Running in Windows WSL and Windows Docker

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
```

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

