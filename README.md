# 📦 Go JWT API

API RESTful simples em Go com autenticação via JWT e banco de dados PostgreSQL.

---

## 🚀 Funcionalidades

- Registro de usuários
- Login com geração de token JWT
- Listagem, atualização e remoção de usuários (rotas protegidas)
- Hash seguro de senhas com bcrypt
- Validação de token via middleware

---

## 📁 Como rodar o projeto localmente

### 1. Clone o repositório

```bash
git clone https://github.com/wendelmatheus/autenticacao-golang.git
cd autenticacao-golang
```

### 2. Crie um arquivo `.env.development` na raiz do projeto

Com as seguintes variáveis:

```env
DB_USER=usuario
DB_PASSWORD=senhaDoBanco
DB_NAME=nomeDoBanco
DB_HOST=localhost
DB_PORT=5432

JWT_SECRET=suaChaveSecretaJWT
```

> 📝 Substitua os valores conforme seu ambiente local.

---

### 3. Crie o banco PostgreSQL

Certifique-se de que o banco com os dados acima existe. Você pode criar assim (no terminal do Postgres):

```sql
CREATE DATABASE nomeDoBanco;
```

E certifique-se de ter a tabela:

```sql
CREATE TABLE usuarios (
  id SERIAL PRIMARY KEY,
  nome TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  senha TEXT NOT NULL
);
```

---

### 4. Instale as dependências

Se ainda não tiver, instale o Go:
https://go.dev/doc/install

Depois, no terminal:

```bash
go mod tidy
```

---

### 5. Rode o servidor

```bash
go run main.go
```

O servidor provavelmente rodará em:  
📍 `http://localhost:8080`

---

## 📫 Rotas disponíveis

### 🔓 Públicas

- `POST /register` – Criação de usuário
- `POST /login` – Login e retorno do token JWT

### 🔐 Protegidas (requer token Bearer)

- `GET /usuarios` – Listar todos os usuários
- `GET /usuarios/{id}` – Buscar um usuário por ID
- `PUT /usuarios/{id}` – Atualizar um usuário
- `DELETE /usuarios/{id}` – Deletar um usuário

> No Postman, insira o token JWT no header:
> 
> ```
> Authorization: Bearer SEU_TOKEN_AQUI
> ```
