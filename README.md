# Golang POC - Senior Side Car
Esse projeto tem o objetivo de levantar quais foram as facilidades/dificuldades de criar uma aplicação REST em Golang. 

## Sobre
**Versão da linguagem**: 1.18  
**IDE utilizada**: VSCode com a extensão GO. 

```Brinquei um pouco na versão EAP do Goland disponibilizada pela JetBrains com 30 dias de teste, de fato é mais prático, mas com a extensão do GO no VSCode já quebra o galho.```

## Estrutura
APIs disponibilizadas:  
* **/getUsers**: Traz uma lista de usuários.
* **/deserialize**: Chama a API /getUsers e retorna o resultado desserializado como text/plain formatado pelo GO.

## Bibliotecas
Não foram utilizadas bibliotecas externas.

## Facilidades

- Simples de montar uma aplicação rest:    
  - não precisa utilizar nenhuma biblioteca externa (embora existam várias que facilitam em vários aspectos).
  - nos primeiros minutos de contato com a linguagem já está pronta.
- Passa a sensação de ser bem mais produtivo do que Java (que é a base que tenho pra comparar).
- Encontrei bastante artigos/vídeos na internet e tópicos no `stackoverflow` sobre soluções simples que precisei utilizar.

## Dificuldades
- Pointers (`&` e `*` antes das variáveis) — até então não consegui compreender exatamente quando devo ou não usá-los.
