# 🔐 GoShop Auth Service

The **Auth Service** is a microservice in the GoShop ecosystem responsible for **user authentication and authorization**. It internally integrates with **FusionAuth** for identity management and issues **JWT tokens** to authenticate users across other GoShop services.

---

## 📦 Features

- 🔑 **User Signup**  
  Register new users using FusionAuth via secure API.

- 🔐 **User Login**  
  Authenticate users and issue JWT tokens.

- 🧾 **JWT Authentication**  
  Verify and validate JWTs for protected routes.

- 🔗 **FusionAuth Integration**  
  Uses FusionAuth APIs for secure user management.

- 🌐 **RESTful APIs**  
  Provides clean and secure endpoints for auth operations.

---

## 🧰 Tech Stack

| Component       | Technology    |
|----------------|---------------|
| Language        | Go            |
| Auth Provider   | FusionAuth    |
| Authentication  | JWT (JSON Web Token) |
| Web Framework   | Gin           |
| HTTP Client     | `net/http` or third-party (e.g. resty) |

---

## 🚀 API Endpoints

| Method | Endpoint         | Description              |
|--------|------------------|--------------------------|
| POST   | `/api/v1/signup` | Register a new user      |
| POST   | `/api/v1/login`  | Authenticate user & get JWT |
| GET    | `/api/v1/me`     | Get current user profile (JWT required) |

---

## 🧼 Security Notes

- Never expose FusionAuth API keys in public repos.
- JWT secret should be strong and kept in a secure environment (e.g. Vault, AWS Secrets Manager).

---

## ✨ Author

Made with ❤️ by **Vamsi Rayapati**
