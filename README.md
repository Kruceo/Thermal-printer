### Thermal Printer Web API

#### Description

The Thermal Printer Web API provides an endpoint for printing text on a thermal printer. The thermal printer needs to be connected to the network or the device where the service is running.

#### Endpoint

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
```