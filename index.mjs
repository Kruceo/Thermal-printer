import { ThermalPrinter, PrinterTypes, CharacterSet } from 'node-thermal-printer';
import express from 'express'
import cfg from './config/config.json' assert {type: "json"}
import cors from 'cors'
const app = express()
app.use(express.json())
app.use(cors())

let printer = new ThermalPrinter({
  type: PrinterTypes[cfg.printer.type],
  interface: cfg.printer.interface,
  characterSet: CharacterSet[cfg.printer.characterSet]
});
printer.println("rodando")
printer.execute()
let usingPrint = false

async function checkPrinting() {
  const promise = new Promise((res, rej) => {
    if (!usingPrint) {
      res()
    }
    else {
      const interval = setInterval(() => {
        console.log('checking')
        if (!usingPrint) {
          res()
          clearInterval(interval)
        }
      }, 500)
    }
  })

  await promise
}

let count = 0

app.get('/get/width', async (req, res) => {
  res.json({ width: printer.getWidth() })
})

app.post('/query', async (req, res) => {
  console.log("query " + count + ' from ' + req.hostname)
  count++
  const { query } = req.body

  console.log(req.body)

  await checkPrinting()
  usingPrint = true
  for (const q of query) {
    switch (q[0]) {
      case "println":
        printer.println(q[1])
        break;
      case "cut":
        printer.cut()
        break;
      case "center":
        printer.alignCenter()
        break;
      case "left":
        printer.alignLeft()
        break;
      case "right":
        printer.alignRight()
        break;
      case "qrcode":
        printer.printQR(q[1])
        break;
      case "code128":
        printer.code128(q[1])
        break;


      default:
        break;
    }
  }

  await printer.execute()
  printer.clear()
  usingPrint = false

  res.json({ error: false })
})

app.listen(cfg.port, () => console.log(`running at ${cfg.port}`))
