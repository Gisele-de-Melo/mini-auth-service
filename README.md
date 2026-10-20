# 🔐 Mini Auth Service
Um serviço de autenticação moderno, rápido e seguro, desenvolvido em **Go**, usando **JWT**, **bcrypt**, **SQLite**, **Chi Router** e arquitetura limpa.  
Ideal para portfólios, estudos ou como base para sistemas maiores.

---

## 🚀 Tecnologias Utilizadas

- **Go 1.26.3
- **Chi Router** (HTTP router leve e poderoso)
- **JWT (HS256)** para autenticação
- **bcrypt** para hashing de senha
- **SQLite (modernc.org/sqlite)** sem CGO
- **Middleware de Rate Limit**
- **Arquitetura limpa (handlers, services, repository, security)**

---

## 📌 Funcionalidades

- Registro de usuário
- Login com geração de JWT
- Rota protegida `/me`
- Hash seguro de senha (bcrypt)
- Validação de token JWT
- Rate limiting por IP
- Estrutura modular e escalável

---

## 🏗 Arquitetura

```
mini-auth-service/
│
├── cmd/api/main.go
│
├── internal/
│   ├── config/          # Configurações
│   ├── http/
│   │   ├── handlers/    # Handlers HTTP
│   │   ├── middleware/  # JWT + Rate Limit
│   │   └── router.go    # Rotas
│   ├── model/           # Modelos
│   ├── repository/      # Persistência (SQLite)
│   ├── security/        # JWT + bcrypt
│   └── service/         # Regras de negócio
│
└── pkg/logger/          # Logger simples
```

---

## ⚙️ Como Rodar

### 1️⃣ Clone o repositório

```bash
git clone https://github.com/seu-usuario/mini-auth-service.git
cd mini-auth-service
```

### 2️⃣ Instale dependências

```
go mod tidy
```

### 3️⃣ Execute o servidor

```
go run cmd/api/main.go
```

Servidor rodará em:
```
http://localhost:8080
```

## 🧪 Testando com curl

### Registrar usuário
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"email":"gisele@example.com","password":"123456"}'
```

### Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"gisele@example.com","password":"123456"}'
```
Resposta:
```json
{"token":"<JWT_AQUI>"}
```

### Rota protegida /me
```bash
curl http://localhost:8080/me \
  -H "Authorization: Bearer <JWT_AQUI>"
```

## 📚 Endpoints

### 🔓 Público
| Método | Rota | Descrição |
| --- | --- | --- |
| POST | /register | Cria um novo usuário |
| POST | /login | Autentica e retorna JWT |

### 🔐 Protegido (JWT obrigatório)
| Método | Rota | Descrição |
| --- | --- | --- |
| GET | /me | Retorna dados do usuário autenticado |

## 🔒 Segurança
- Hash de senha com bcrypt
- Tokens JWT com HS256
- Expiração automática de token (2h)
- Rate limit por IP (20 req/min)
- Middleware de proteção para rotas privada

## 📄 Licença
Este projeto é open-source e pode ser usado livremente.  

## 👩‍💻 Autora
Gisele de Melo  
Desenvolvedora Backend • Go • APIs • Arquitetura Limpa
Projeto desenvolvido com caráter de estudo.
