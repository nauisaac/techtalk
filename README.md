# Tech Talk - Marvel API MCP Server

Este projeto implementa um servidor MCP (Model Context Protocol) que integra com a Marvel API.

## Configuração

1. Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```env
# Marvel API Configuration
# Get your keys at https://developer.marvel.com/
MARVEL_BASE_URL=https://gateway.marvel.com
MARVEL_PUBLIC_KEY=your_public_key_here
MARVEL_PRIVATE_KEY=your_private_key_here
```

2. Obtenha suas chaves da API Marvel em: https://developer.marvel.com/

3. Execute o projeto:

```bash
go run main.go
```

## Funcionalidades

- **get-character**: Busca informações de personagens da Marvel por nome
- **calculate**: Realiza operações matemáticas básicas

## Correções Implementadas

- ✅ Hash MD5 obrigatório para autenticação da Marvel API
- ✅ Encoding correto de parâmetros na URL
- ✅ Melhor tratamento de erros da API
- ✅ Documentação de configuração

