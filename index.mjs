import { ThermalPrinter, PrinterTypes, CharacterSet, BreakLine } from 'node-thermal-printer';
import express from 'express'

const port = 1000

const app = express()
app.use(express.json())

let printer = new ThermalPrinter({
  type: PrinterTypes.EPSON,                                  // Printer type: 'star' or 'epson'
  interface: '/dev/usb/lp0',
  characterSet: CharacterSet.SLOVENIA                        // Printer interface
});

let usingPrint = false

async function checkPrinting (){
  const promise = new Promise((res,rej)=>{
    if(!usingPrint){
      res()
    }
    else {
      const interval = setInterval(()=>{
        if(!usingPrint){
          res()
          clearInterval(interval)
        }
      },500)
    }
  })

  await promise
}

let count = 0
app.post('/query',async (req,res)=>{
  console.log("query " + count + ' from ' + req.hostname)
  count ++
  const {query} = req.body
  
  await checkPrinting()
  usingPrint = true
  for (const q of query) {
    switch (q[0]) {
      case "println":
          await new Promise((res)=>setTimeout(()=>res(),1000))
          printer.println(q[1])
          await printer.execute()
          printer.clear()
          
        break;
        case "cut":
          printer.cut()
        break;
    
      default:
        break;
    }
  }

  await printer.execute()
  printer.clear()
  usingPrint = false

  res.json({error:false})
})

app.listen(port,()=>console.log(`running at ${port}`))
