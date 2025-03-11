# GenPass - Gerador de Senhas

Um aplicativo simples em Golang para gerar senhas aleatórias com opções personalizáveis.

## Funcionalidades

- Gera senhas aleatórias com opções personalizáveis
- Permite especificar quais tipos de caracteres incluir (números, letras, caracteres especiais)
- Permite definir o comprimento da senha

## Opções

O gerador de senhas aceita os seguintes argumentos de linha de comando:

- `-n` - Incluir números na senha
- `-l` - Incluir letras na senha (minúsculas e maiúsculas)
- `-s` - Incluir caracteres especiais na senha
- `-c [número]` - Definir o comprimento da senha (padrão: 12)

Se nenhum argumento for especificado, o gerador criará uma senha de 12 caracteres incluindo todos os tipos (números, letras e caracteres especiais).

## Exemplos de Uso

```bash
# Senha padrão (12 caracteres com todos os tipos)
./genpass

# Senha de 16 caracteres com todos os tipos
./genpass -c 16

# Senha de 8 caracteres apenas com números
./genpass -n -c 8

# Senha com números e letras (sem caracteres especiais)
./genpass -n -l

# Senha de 20 caracteres com letras e caracteres especiais
./genpass -l -s -c 20
```

## Instalação

Existem várias maneiras de instalar e usar o GenPass:

### Binários pré-compilados

Você pode baixar os binários compilados na [página de releases](../../releases) para seu sistema operacional.

### Docker (via GitHub Packages)

```bash
# Puxar a imagem
docker pull ghcr.io/nilsonvieira/devops-go-genpass:latest

# Executar o GenPass (com os parâmetros desejados)
docker run ghcr.io/nilsonvieira/devops-go-genpass -l -n -c 16
```

### Go Install

```bash
# Instalar via Go
go install github.com/nilsonvieira/devops-go-genpass@latest

# Usar o comando
genpass -l -n -c 16
```

### Compilação Manual

1. Clone o repositório:
   ```bash
   git clone https://github.com/nilsonvieira/devops-go-genpass.git
   cd devops-go-genpass
   ```

2. Compile o aplicativo:
   ```bash
   go build -o genpass .
   ```

3. Execute o gerador de senhas:
   ```bash
   ./genpass
   ```

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).