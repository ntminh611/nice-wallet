<div align="center">

# 🔐 Nice Wallet Generator

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge)](CONTRIBUTING.md)

*Generate vanity wallets with beautiful custom address patterns for all EVM-compatible blockchains - Create wallets ending with 888, 999, 000 or any pattern you want!* 🎯

<p align="center">
  <img src="https://img.shields.io/badge/Ethereum-3C3C3D?style=for-the-badge&logo=Ethereum&logoColor=white">
  <img src="https://img.shields.io/badge/BSC-F0B90B?style=for-the-badge&logo=binance&logoColor=white">
  <img src="https://img.shields.io/badge/Arbitrum-2D374B?style=for-the-badge&logo=arbitrum&logoColor=white">
  <img src="https://img.shields.io/badge/Base-0052FF?style=for-the-badge&logo=base&logoColor=white">
  <img src="https://img.shields.io/badge/Polygon-8247E5?style=for-the-badge&logo=polygon&logoColor=white"> ...
</p>

</div>

---

## ✨ Features

- 🎯 Generate wallets with custom patterns
- 🚀 High-performance Go implementation
- 💾 Auto-save results to file
- 🔄 Cross-platform compatibility
- 🔒 Secure key generation

## 🎯 Pattern Examples

The pattern matching system uses the following symbols:

| Symbol | Meaning |
|--------|---------|
| `*` | Any sequence of characters |
| `x` | Single character |

### 📝 Example Patterns
```yaml
*999         # Address ending with "999"
0x1xxx       # Address starting with "0x1" followed by any 3 characters
*123*456     # Address containing "123" followed by "456"
*0000        # Address ending with "0000"
*abc*def     # Address containing "abc" followed by "def"
0xabc*       # Address starting with "0xabc"
*dead        # Address ending with "deed"
*1111*2222   # Address containing "1111" followed by "2222"
```

## 🚀 Prerequisites

### Install Go

<details>
<summary>🪟 Windows</summary>

1. Download Go installer from [official Go website](https://golang.org/dl/)
2. Run the installer (e.g., `go1.21.windows-amd64.msi`)
3. Add Go to your PATH environment variable
4. Verify installation:
```bash
go version
```
</details>

<details>
<summary>🍎 macOS</summary>

Using Homebrew:
```bash
brew install go
```
</details>

<details>
<summary>🐧 Linux (Ubuntu/Debian)</summary>

```bash
sudo apt update
sudo apt install golang-go
```
</details>

## 📥 Installation

1. Clone the repository:
```bash
git clone https://github.com/ntminh611/nice-wallet
cd nice-wallet
```

2. Install dependencies:
```bash
go mod download
```

## ⚙️ Configuration

Create or modify `structs.txt` file with your desired patterns:

```text
*999
0x1xxx
*0000
```

## 🛠️ Build

### Build for current platform
```bash
go build -o nice-wallet
```

### 🌍 Cross-platform builds

<details>
<summary>Build for Windows 🪟</summary>

```bash
GOOS=windows GOARCH=amd64 go build -o nice-wallet.exe
```
</details>

<details>
<summary>Build for macOS 🍎</summary>

```bash
GOOS=darwin GOARCH=amd64 go build -o nice-wallet-mac
```
</details>

<details>
<summary>Build for Linux 🐧</summary>

```bash
GOOS=linux GOARCH=amd64 go build -o nice-wallet-linux
```
</details>

## 🚀 Usage

1. Run the executable:

```bash
# Windows
nice-wallet.exe

# macOS/Linux
./nice-wallet
```

2. The program will generate wallets until it finds matches for all patterns
3. Results are saved in `results.txt`

## 🔐 Security Notice

> [!WARNING]
> **Keep your private keys safe!** Never share them with anyone. Generated wallets are real cryptocurrency wallets.

## 🤝 Contributing

Contributions are welcome! Feel free to submit issues and pull requests.

## 📄 License

MIT License © ntminh611
