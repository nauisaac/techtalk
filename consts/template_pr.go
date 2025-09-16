package consts

var TemplatePR = `# %s

## 📝 Descrição
%s

## 🎯 Tipo de Mudança
- [ ] 🐛 Correção de bug (mudança que corrige um problema)
- [ ] ✨ Nova funcionalidade (mudança que adiciona funcionalidade)
- [ ] 💥 Mudança que quebra compatibilidade (correção ou funcionalidade que causa quebra)
- [ ] 📚 Atualização de documentação
- [ ] 🎨 Melhoria de código (formatação, renomeação de variáveis, etc.)
- [ ] ⚡ Melhoria de performance
- [ ] ✅ Adição ou atualização de testes
- [ ] 🔧 Mudanças de configuração ou build

## 🔄 Mudanças Realizadas
<!-- Descreva as principais mudanças implementadas baseando-se no git diff -->

## 🧪 Como Testar
<!-- Descreva os passos para testar as mudanças -->
1. 
2. 
3. 

## 📸 Screenshots/GIFs (se aplicável)
<!-- Adicione capturas de tela ou GIFs demonstrando as mudanças visuais -->

## 📋 Checklist
- [ ] Meu código segue os padrões de estilo do projeto
- [ ] Realizei uma auto-revisão do meu código
- [ ] Comentei meu código, especialmente em áreas difíceis de entender
- [ ] Fiz mudanças correspondentes na documentação
- [ ] Minhas mudanças não geram novos warnings
- [ ] Adicionei testes que provam que minha correção é efetiva ou que minha funcionalidade funciona
- [ ] Testes unitários novos e existentes passam localmente com minhas mudanças

## 🔄 Plano de Rollback
<!-- Descreva como reverter as mudanças se necessário -->

---

## 📚 Guia de Criação de PR

### 🏷️ Formato do Título
- **Padrão de branch**: developer/CRD-XXX-description
- **Título do PR**: [CRD-XXX] - Description
- **Sem numeração específica**: Use CRD-000


### 🚀 Criação via GitHub CLI
` + "```bash" + `
gh pr create --base main --head your-branch --title "[CRD-XXX] - Description" --body-file pr_template.md
` + "```" + `

### ⚠️ Lembrete Importante
- **NÃO** cole apenas o template em branco
- **PREENCHA** todas as seções relevantes baseando-se no que foi implementado
- **ANALISE** o git diff para descrever adequadamente as mudanças
- **TESTE** suas mudanças antes de abrir o PR

`
