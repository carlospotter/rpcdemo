# rpcdemo

Demo de comunicação utilizando o protocolo de comunicação RPC (Remote Procedure Call) e gRPC, framework RPC desenvolvido pelo Google.

## RPC

RPC (Remote Procedure Call) é um protocolo de comunicação que permite que uma aplicação execute um trecho de código de outra aplicação, fazendo com que processos ou sistemas distintos se comuniquem pela rede, abstraindo as complexidades de comunicação entre cliente e servidor.

Golang permite a criação de um cliente e um servidor RPC utilizando a biblioteca padrão ["net/rpc"](https://pkg.go.dev/net/rpc). Um ponto negativo dessa abordagem é que tanto o server como o client precisam ser desenvolvidos utilizando Go como linguagem.

## gRPC

gRPC (Google Remote Procedure Call) é um framework RPC desenvolvido pelo Google. Além de vantagens de performance, ele também possui suporte para múltiplas linguagens, permitindo que o client e o server sejam desenvolvidos utilizando linguagens de programação distintas.

Para isso, o protocolo deve ser definido em um arquivo `.proto`, onde são definidas as propriedades da request, response, métodos expostos pelo protocolo. Utilizando [ferramentas específicas](https://grpc.io/), esse arquivo `.proto` é compilado e o código para exposição dos métodos e criação de structs é gerado automaticamente, faltando apenas a criação do client e do server utilizando o package [grpc](https://pkg.go.dev/google.golang.org/grpc).