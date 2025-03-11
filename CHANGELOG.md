# Changelog

Todas as mudanças notáveis neste projeto serão documentadas neste arquivo.

O formato é baseado em [Keep a Changelog](https://keepachangelog.com/pt-BR/1.0.0/),
e este projeto adere ao [Versionamento Semântico](https://semver.org/lang/pt-BR/).

## [v1.0.1] - 2025-03-11

### Corrigido
- Atualização das ações do GitHub Actions (de v3 para v4) para resolver problemas de compatibilidade
- Correção no processo de inicialização do módulo Go para evitar erros quando o arquivo go.mod já existe
- Otimização do workflow para melhor desempenho no CI/CD

## [v1.0.0] - 2025-03-11

### Adicionado
- Lançamento inicial do GenPass
- Implementação do gerador de senhas em Go
- Opções de linha de comando (-n, -l, -s, -c)
- Suporte para gerar senhas com números, letras e caracteres especiais
- Configuração da pipeline de CI/CD com GitHub Actions
- Disponibilização de binários pré-compilados para Linux, Windows e macOS
