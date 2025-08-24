# ai-deploy-one-click

CLI command to deploy your AI chatbot in one click on Vercel.

## Pre-requisites

In order to use **ai-deploy-one-click**, you need to have the following tools installed:

- [Vercel CLI](https://vercel.com/docs/cli)
- [GitHub CLI](https://cli.github.com)
- [Git](https://git-scm.com)

For what concerns `git`, you need to have push permission without the need to pass username and authentication token.

## Install

In order to install **ai-deploy-one-click** there are three ways:

1. Using `go` (`go` 1.23+ required):

```bash
go install github.com/AstraBert/ai-deploy-one-click
```

2. Using `npm`:

```bash
npm install @cle-does-things/ai-deploy-one-click
```

3. Downloading the executable from the [releases page](https://github.com/AstraBert/ai-deploy-one-click/releases): you can download it directly from the GitHub repository or, if you do not want to leave your terminal, you can use `curl`:

```bash
curl -L -o ai-deploy-one-click https://github.com/AstraBert/ai-deploy-one-click/releases/download/<version>/ai-deploy-one-click_<version>_<OS>_<processor>.tar.gz ## e.g. https://github.com/AstraBert/ai-deploy-one-click/releases/download/0.1.1/ai-deploy-one-click_0.1.1_darwin_amd64.tar.gz

# make sure the downloaded binary is executable (not needed for Windows)
chmod +x ai-deploy-one-click
```

In this last case, be careful to specify your OS (supported: linux, windows, macos) and your processor type (supported: amd, arm).

Install with:

```bash
npm install @cle-does-things/ai-deploy-one-click
```

## Run

You can run **ai-deploy-one-click** simply like this from your terminal:

```bash
ai-deploy-one-click
```

This will start a terminal user interface in which you will be asked to enter the details of your AI application (name, description, URL, GitHub source, AI model, API key...).

These information will be used to create a GitHub repository from [a template AI chatbot](https://github.com/AstraBert/ai-deploy-one-click-ui), clone the repository locally and configure it based on the information you provided: once the application configured, it will be connected to Vercel, your API key will be added to the production environment and the local changes will be pushed to GitHub, starting a deployment on Vercel. In less than 5 minutes you will be able to see your AI Chatbot up and running online!

## Contributing

We welcome contributions! Please read our [Contributing Guide](./CONTRIBUTING.md) to get started.

## License

This project is licensed under the [MIT License](./LICENSE)
