import usb.core
import usb.util

dev = usb.core.find(idVendor=0x04b8, idProduct=0x0e27)

# Verifique se o dispositivo foi encontrado
if dev is None:
    raise ValueError('Dispositivo USB não encontrado')



if(dev.is_kernel_driver_active(0)):
    dev.detach_kernel_driver(0)
    print("detached kernel driver 0")
    # exit()

dev.set_configuration()

usb.util.claim_interface(dev, 0)
print ("claimed device")

# Conecte ao dispositivo


# Envie uma string para a porta USB
dev.write(0x01, "Sua string aqui\n",timeout=1000)

# Feche a conexão USB
usb.util.dispose_resources(dev)
