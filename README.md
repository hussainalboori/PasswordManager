# JassPass — Modern & Secure Password Manager

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

**JassPass** is a high-performance, self-contained password manager built using **Go (Golang)** and a pure-Go **SQLite** database. It features AES-256 credential encryption, a dark glassmorphic user interface, real-time vault searching, 1-click clipboard actions, password strength evaluation, and Vercel serverless deployment support.

---

## Key Features

- **Fortress-Level Encryption**: Passwords are encrypted using **AES-256** prior to database storage, paired with per-user unique encryption keys.
- **Glassmorphic Dark Interface**: Designed with CSS glassmorphism (`backdrop-filter`), glowing color accents, responsive card grids, and smooth micro-interactions.
- **Real-Time Vault Search**: Filter stored website credentials instantly by domain or username as you type.
- **1-Click Copy & Toast Notifications**: Copy usernames or decrypted passwords to clipboard with single-click actions and animated toast confirmation alerts.
- **Smart Password Generator**: Custom entropy generator with adjustable length (8–32 characters), symbol, and number toggles.
- **Real-Time Password Strength Meter**: Live strength analyzer providing instant visual feedback (Weak, Good, Strong).
- **Embedded Binary Assets**: HTML templates and CSS styles are embedded directly into the Go executable via `//go:embed` for zero-dependency deployment.
- **Vercel Serverless Ready**: Native deployment support for Vercel using pure-Go CGO-free SQLite (`modernc.org/sqlite`) and `/tmp` database path abstraction.

---

## Tech Stack

- **Backend**: Go (Golang)
- **Database**: SQLite (`modernc.org/sqlite` — 100% CGO-free)
- **Frontend**: HTML5, Modern Vanilla CSS (Glassmorphism), Vanilla JavaScript
- **Asset Bundling**: Native Go `embed.FS`
- **Deployment**: Vercel Serverless / Local Go binary

---

## Getting Started

### Prerequisites
- [Go 1.21+](https://go.dev/dl/) installed on your system.

### Installation

Clone the repository:
```bash
git clone https://github.com/hussainalboori/PasswordManager.git
cd PasswordManager
```

---

## Running Locally

To run the application locally:
```bash
go run .
```

The application automatically checks port `8080`. If port `8080` is occupied, it gracefully binds to the next available port (`8081`, `8082`, etc.):

```text
Database file 'data.db' already exists.
2026/09/01 13:38:45 Connect to our website through http://localhost:8081
```

### Specifying a Custom Port
You can specify a custom port using the `PORT` environment variable:
```bash
PORT=9000 go run .
```

---

## Deploying to Vercel

JassPass is pre-configured for Vercel serverless deployment (`api/index.go` and `vercel.json`).

### 1-Click CLI Deployment
Install the Vercel CLI and deploy to production:
```bash
npx vercel --prod
```

Because the database engine uses `modernc.org/sqlite` (pure Go without CGO requirements) and templates are embedded into the binary via `//go:embed`, the project compiles and runs seamlessly on serverless platforms.

---

## Project Architecture

```text
PasswordManager/
├── api/
│   └── index.go          # Vercel serverless entrypoint handler
├── data/
│   ├── Database.go       # SQLite connection manager & table initializers
│   ├── EncryptDecryptionPassword.go # AES-256 encryption/decryption logic
│   ├── AddNewPassword.go # Insert password record
│   ├── getuserbyemail.go # User authentication query logic
│   └── ...
├── handler/
│   ├── Dashbored.go      # Vault dashboard route handler
│   ├── login.go          # Login & session management
│   ├── sigup.go          # User registration
│   ├── new.go            # Add password & generator handler
│   └── template_helper.go# Embedded template renderer
├── static/
│   ├── css/style.css     # Glassmorphic design system
│   └── embed.go          # Embedded static assets (embed.FS)
├── templates/
│   ├── index.html        # Landing page
│   ├── login.html        # Authentication page
│   ├── signup.html       # Sign up page
│   ├── dashboard.html    # Vault grid & search dashboard
│   ├── new.html          # Add credential & generator
│   └── embed.go          # Embedded HTML templates (embed.FS)
├── vercel.json           # Vercel deployment routes config
└── server.go             # Local Go web server entrypoint
```

---

## Security Considerations & Disclaimer

This password manager is designed for educational and demonstration purposes. While passwords are encrypted using AES-256:
- Always use strong master passwords.
- For production multi-user deployments, master key derivation (e.g. PBKDF2/Argon2) and secure KMS key management should be implemented.

---

## License

This project is open source and available under the [MIT License](file:///Users/abojass/Documents/GitHub/PasswordManager/LICENSE):

```text
MIT License

Copyright (c) 2024 - 2026 Hussain Alboori

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```