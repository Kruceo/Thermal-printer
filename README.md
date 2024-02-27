### API Web de Impressora Térmica

#### Descrição

A API Web de Impressora Térmica fornece um endpoint que permite imprimir textos em uma impressora térmica. A impressora térmica precisa estar conectada à rede ou ao dispositivo no qual o serviço estiver sendo executado.

#### Endpoint

- **URL**: `POST /query`
- **Descrição**: Esse endpoint recebe uma lista de comandos e textos a serem impressos e os envia para a impressora térmica. Os comandos são fornecidos no corpo da solicitação como uma lista de strings.

#### Corpo da Solicitação (Exemplo)

```json
{
  "query": [
    [ "center" ],
    [ "println", "Transacao 47" ],
    [ "left" ],
    [ "println", "------------------------------------------------" ],
    [ "println", "Tipo: Entrada" ],
    [ "println", "Bote: SSC Tuatara" ],
    [ "println", "Fornecedor: Pensilvania" ],
    [ "code128", "https://kruceo.com"]
    [ "qrcode" , "https://kruceo.com"]
    [ "cut" ]
  ]
}
```
