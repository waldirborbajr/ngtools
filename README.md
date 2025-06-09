# NGTools

<p align="center">
  <img width="256" height="256" src="./assets/ngtools-logo.png" />
</p>

[![lint](https://github.com/waldirborbajr/ngtools/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/waldirborbajr/ngtools/actions/workflows/golangci-lint.yml)

`tl;dr:` NGTools is a CLI (Command Line Interface), that wraps `NGrok` to start and stop and returns HTTPS.

---

## 📁 Project Structure

```
ngtools/
├── cmd/
│   ├── addr.go         # Command to display ngrok public URL
│   ├── root.go         # CLI definition and command routing
│   ├── start.go        # Command to start ngrok
│   ├── stop.go         # Command to stop ngrok
│   └── version.go      # Command to display ngrok version
├── internal/
│   ├── getngrokurl/    # Fetches ngrok public URL via local API
│   ├── hascurl/        # Checks if curl is installed
│   ├── hasnohup/       # Manages nohup.out file
│   ├── killprocess/    # Kills ngrok process by PID
│   ├── listprocess/    # Checks if ngrok is running
│   ├── secureexec/     # Secure execution of external commands
│   ├── showerror/      # Displays formatted error messages
│   ├── startngrok/     # Starts ngrok and saves PID
│   └── verifyos/       # Checks operating system
├── main.go             # Application entry point
├── go.mod              # Go dependencies management
└── README.md           # This file
```

---

## 🚀 How to Run

### 1. Prerequisites

- Go 1.20+ installed
- ngrok installed and available in your PATH
- Linux (Windows is not supported)

### 2. Installation

Clone the repository and install dependencies:

```sh
git clone https://github.com/youruser/ngtools.git
cd ngtools
go mod tidy
```

### 3. Build

```sh
go build -o ngtools
```

### 4. Usage

```sh
./ngtools <command>
```

#### Available commands:

- `start`   – Starts ngrok and saves the PID
- `stop`    – Stops ngrok using the saved PID
- `addr`    – Displays the ngrok public URL
- `version` – Shows the installed ngrok version

Example:

```sh
./ngtools start
./ngtools addr
./ngtools stop
```

---

## 🛡️ Security and Best Practices

- Does not execute shell commands with user input.
- Restricts permissions for temporary files (0600).
- Only kills the ngrok process started by the tool.
- Not supported on Windows.

---

## 🧪 Testing

To run automated tests:

```sh
go test ./...
```

---

## 📄 License
