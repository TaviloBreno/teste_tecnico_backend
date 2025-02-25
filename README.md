# Gerenciador de Tarefas

Este projeto é um **Gerenciador de Tarefas** desenvolvido em **Flutter**, estruturado com boas práticas, princípios **SOLID** e **Clean Code**, para demonstrar uma arquitetura robusta e escalável.

## 🎯 Objetivo

Criar uma aplicação para gerenciamento de tarefas que permita adicionar, listar, atualizar e excluir tarefas, utilizando uma arquitetura bem definida e organizada.

---

## 🛠️ Tecnologias Utilizadas

- **Flutter**: Framework para desenvolvimento multiplataforma.
- **Dart**: Linguagem de programação.
- **Bloc/Cubit**: Gerenciamento de estado.
- **SOLID**: Princípios de design para código limpo e modular.
- **Clean Code**: Boas práticas para manter o código legível e sustentável.

---

## 📂 Estrutura do Projeto

```
lib/
├── core/
│   ├── constants/          # Constantes globais
│   ├── exceptions/         # Exceções personalizadas
│   ├── utils/              # Utilitários e helpers
│   └── theme/              # Temas da aplicação
├── data/
│   ├── datasources/        # Fontes de dados (local ou remoto)
│   │   └── local/          # Implementação do DataSource Local
│   ├── models/             # Modelos de dados
│   └── repositories/       # Repositórios concretos
├── domain/
│   ├── entities/           # Entidades do domínio
│   ├── repositories/       # Contratos dos repositórios
│   └── usecases/           # Casos de uso
├── presentation/
│   ├── pages/              # Páginas da interface do usuário
│   ├── widgets/            # Componentes reutilizáveis
│   ├── cubit/              # Gerenciamento de estado
│   └── routes/             # Configuração de rotas
├── main.dart                # Ponto de entrada da aplicação
```

---

## 🔑 Funcionalidades

- **Listar tarefas**: Exibe todas as tarefas salvas.
- **Adicionar tarefas**: Permite adicionar uma nova tarefa.
- **Editar tarefas**: Atualiza informações de uma tarefa existente.
- **Excluir tarefas**: Remove uma tarefa específica.
- **Marcar como concluída**: Define o status de conclusão de uma tarefa.

---

## 🚀 Como Executar

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/seu-usuario/gerenciador-de-tarefas.git
   cd gerenciador-de-tarefas
   ```

2. **Instale as dependências:**

   ```bash
   flutter pub get
   ```

3. **Execute o projeto:**

   ```bash
   flutter run
   ```

---

## 🛡️ Arquitetura

O projeto segue uma arquitetura **limpa** com camadas bem definidas:

- **Core**: Contém elementos genéricos compartilhados por toda a aplicação.
- **Data**: Responsável pela implementação de fontes de dados, como APIs ou armazenamento local.
- **Domain**: Contém as regras de negócio, incluindo entidades, contratos de repositórios e casos de uso.
- **Presentation**: Contém a interface do usuário, gerenciamento de estado e rotas.

---

## 📋 Princípios Aplicados

1. **Single Responsibility Principle (SRP)**: Cada classe tem uma única responsabilidade.
2. **Open/Closed Principle (OCP)**: Classes estão abertas para extensão, mas fechadas para modificação.
3. **Liskov Substitution Principle (LSP)**: Subclasses podem substituir suas superclasses sem alterar o comportamento esperado.
4. **Interface Segregation Principle (ISP)**: Interfaces são específicas e evitam dependências desnecessárias.
5. **Dependency Inversion Principle (DIP)**: Dependência de abstrações, não de implementações.

---

## 🤝 Contribuições

Contribuições são bem-vindas! Siga os passos abaixo para colaborar:

1. Faça um fork do repositório.
2. Crie uma nova branch:
   ```bash
   git checkout -b feature/sua-feature
   ```
3. Commit suas alterações:
   ```bash
   git commit -m "Adiciona nova funcionalidade"
   ```
4. Envie para o seu fork:
   ```bash
   git push origin feature/sua-feature
   ```
5. Abra um Pull Request.

---

## 📄 Licença

Este projeto está licenciado sob a licença MIT. Consulte o arquivo `LICENSE` para mais informações.

✨ Feito com Flutter e Clean Code para um código mais limpo, modular e sustentável! 🚀