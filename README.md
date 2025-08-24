# ai-deploy-one-click

**Deploy your AI chatbot to production in under 5 minutes with a single command.**

A powerful CLI tool that automates the entire deployment process: from repository creation to live deployment on Vercel.

## 📋 Prerequisites

Before using **ai-deploy-one-click**, ensure you have the following tools installed and configured:

| Tool           | Purpose               | Installation Guide                     |
| -------------- | --------------------- | -------------------------------------- |
| **Vercel CLI** | Deployment platform   | [Install](https://vercel.com/docs/cli) |
| **GitHub CLI** | Repository management | [Install](https://cli.github.com)      |
| **Git**        | Version control       | [Install](https://git-scm.com)         |

> **Note**: Git should be configured with push permissions without requiring manual username/token input.

## 📦 Installation

Choose your preferred installation method:

### Option 1: Go Install

```bash
# Requires Go 1.23+
go install github.com/AstraBert/ai-deploy-one-click
```

### Option 2: NPM

```bash
npm install @cle-does-things/ai-deploy-one-click
```

### Option 3: Direct Download

Download the executable from our [releases page](https://github.com/AstraBert/ai-deploy-one-click/releases):

```bash
# Using curl (replace placeholders with your values)
curl -L -o ai-deploy-one-click \
  https://github.com/AstraBert/ai-deploy-one-click/releases/download/<version>/ai-deploy-one-click_<version>_<OS>_<processor>.tar.gz

# Make executable (Unix/Linux/macOS only)
chmod +x ai-deploy-one-click
```

**Supported platforms:**

- **OS**: Linux, Windows, macOS
- **Processors**: AMD64, ARM64

## 🎮 Usage

Launch the interactive setup with a single command:

```bash
ai-deploy-one-click
```

### 🤖 Supported AI Models

| Provider      | Models Available                                      |
| ------------- | ----------------------------------------------------- |
| **OpenAI**    | GPT 4.1, GPT 4o, GPT 5                                |
| **Anthropic** | Claude Sonnet 3.5, Claude Sonnet 3.7, Claude Sonnet 4 |
| **Google**    | Gemini 2 Flash, Gemini 2.5 Flash, Gemini 2.5 Pro      |

### What happens next?

1. **📝 Configuration**: Enter your app details through an intuitive terminal interface:

   - Application name and description
   - GitHub repository details
   - AI model selection (GPT, Claude, Gemini)
   - System prompts and API keys

2. **🏗️ Repository Setup**: Automatically creates a new GitHub repository from our [template](https://github.com/AstraBert/ai-deploy-one-click-ui)

3. **⚙️ Configuration**: Clones and configures the repository with your specifications

4. **🚀 Deployment**: Connects to Vercel, sets up environment variables, and triggers deployment

5. **✅ Live**: Your AI chatbot is ready and accessible online!

```mermaid
graph LR
    A[Run CLI] --> B[Fill Configuration]
    B --> C[Create GitHub Repo]
    C --> D[Clone & Configure]
    D --> E[Connect Vercel]
    E --> F[Deploy Live]
```

## 🤝 Contributing

We welcome contributions! Please read our [Contributing Guide](./CONTRIBUTING.md) to get started.

## 📄 License

This project is licensed under the [MIT License](./LICENSE).
