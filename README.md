# Cryptography_project
# Project Title: "Secure Password Manager using bcrypt and AES Encryption"

📝 Short Description

This project is a **command-line–based secure password manager** developed using the **Go programming language**.
It allows users to store and retrieve account credentials securely using a **single master password**.

The system applies **cryptographic techniques** such as **bcrypt hashing**, **PBKDF2 key derivation**, and **AES encryption** to ensure that passwords are **never stored in plaintext**. All data is saved locally in an encrypted vault file.

This project is designed for **educational purposes** to demonstrate how cryptography is applied in real-world security systems.
Great question — this is **exactly the right thing to document** 
What you want now is a **clear, honest, step-by-step explanation of the environment, dependencies, and setup fixes**, so that **anyone (lecturer, TA, GitHub reviewer)** can run your project **without the problems you faced**.

## 🛠 Environment, Dependencies, and Setup Notes

This project depends on a correctly configured **Go development environment** and external cryptographic libraries. During development, several configuration steps were required to ensure the project runs successfully and can be pushed to GitHub.

This section documents all requirements and setup steps clearly so that others can reproduce the project without errors.

## 💻 Development Environment

* **Operating System:** Windows 10 / Windows 11
* **IDE:** Visual Studio Code (VS Code)
* **Go Version:** Go **1.22 or later**
* **Version Control:** Git + GitHub

The project was developed and tested using **Go Modules** for dependency management.

## 📦 Required Libraries / Dependencies

### 1️⃣ Go Standard Libraries (Built-in)

These libraries are included automatically with Go and require no installation:

| Package             | Purpose                         |
| ------------------- | ------------------------------- |
| `crypto/aes`        | AES encryption                  |
| `crypto/cipher`     | AES-GCM encryption mode         |
| `crypto/rand`       | Secure random number generation |
| `crypto/sha256`     | Hash function for PBKDF2        |
| `encoding/json`     | Vault file storage              |
| `encoding/base64`   | Encode encrypted data           |
| `os`, `io`, `bufio` | File and input handling         |
| `time`              | Timestamp recording             |


### 2️⃣ External Cryptographic Library

This project **requires an external cryptography library**:

```
golang.org/x/crypto
```

Used components:

| Sub-package | Purpose                           |
| ----------- | --------------------------------- |
| `bcrypt`    | Secure hashing of master password |
| `pbkdf2`    | Password-based key derivation     |

This library **is not included by default** and must be installed using Go Modules.

## 📁 Go Module Configuration

The project uses a **Go module** defined in `go.mod`:

```go
module secure-password-manager

go 1.22

require golang.org/x/crypto v0.45.0
```

Go Modules ensure:

* All dependencies are versioned
* The project builds consistently on different machines

## ⚙️ Installation & Setup

Follow these steps **exactly** to avoid errors.

### Step 1: Verify Go Installation

```bash
go version
```

Output should show **Go 1.22 or later**.

### Step 2: Enable Go Modules (Important on Windows)

```powershell
go env GO111MODULE
```

If it is `off`, enable it:

```powershell
setx GO111MODULE on
```

Restart the terminal after running this command.

### Step 3: Install Dependencies

From the project root directory:

```bash
go mod tidy
```

This command:

* Downloads `golang.org/x/crypto`
* Fixes missing or unused dependencies
* Generates `go.sum`

### Step 4: Build the Project (Optional but Recommended)

```bash
go build .
```

If no error appears, the environment is correctly set up.

## ▶️ Running the Program

⚠️ **Important:**
Do **NOT** run only `main.go`.
This project consists of **multiple files** that must be compiled together.

✅ Correct way:

```bash
go run .
```

or with a command:

```bash
go run . init
go run . add
go run . list
go run . get Gmail
```

❌ Incorrect (causes undefined function errors):

```bash
go run main.go
```


## 📂 Vault File Requirement

* The file `vault.json` is **automatically created** when running:

```bash
go run . init
```

* Users **should not create or edit this file manually**
* All passwords inside are **encrypted**

## ▶️ Usage Examples

### 1️⃣ Initialize the Vault

Creates a new encrypted vault and master password.

```bash
go run . init
```

### 2️⃣ Add a Password Entry

Stores a new encrypted credential.

```bash
go run . add
```

### 3️⃣ List All Stored Entries

Displays all stored passwords after authentication.

```bash
go run . list
```

### 4️⃣ Retrieve a Specific Entry

Gets the password for a specific site.

```bash
go run . get <site>
```

Example:

```bash
go run . get Gmail
```

## 👤 Author

**Name:** Tun Monireaksa
**Course:** Cryptography
**Project Type:** Individual Project

