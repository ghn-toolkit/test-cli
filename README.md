[![Create Go App][repo_logo_img]][repo_url]

# Create Go App CLI

[![Go version][go_version_img]][go_dev_url]

Create a new production-ready project with **backend** (Golang),
**frontend** (JavaScript, TypeScript) and **deploy automation** (Ansible,
Docker) by running only one CLI command.

Focus on **writing your code** and **thinking of the business-logic**! The CLI 
will take care of the rest.

## ⚡️ Quick start

First, [download][go_download_url] and install **Go**. Version `1.20` or 
higher is required.

Installation is done by using the [`go install`][go_install_url] command:

```bash
go install github.com/ghn-toolkit/test-cli/cmd/cga@latest
```
GNU/Linux and macOS users available way to install via
[Homebrew][brew_url]:

```bash
# Tap a new formula:
brew tap ghn-toolkit/tap

# Installation:
brew install ghn-toolkit/tap/cga
```

Let's create a new project via **interactive console UI** (or **CUI** for 
short) in current folder:

```bash
cga create
```

Next, open the generated Ansible inventory file (called `hosts.ini`) and 
fill in the variables according to your server configuration. And you're 
ready to **automatically deploy** this project:

```bash
cga deploy
```

That's all you need to know to start! 🎉

## 📖 Project Wiki

The best way to better explore all the features of the **Create Go App CLI** 
is to read the project [Wiki][repo_wiki_url] and take part in 
[Discussions][repo_discussions_url] and/or [Issues][repo_issues_url]. 

Yes, the most frequently asked questions (FAQ) are also 
[here][repo_wiki_faq_url].

## ⚙️ Commands & Options

### `create`

CLI command for create a new project with the interactive console UI.

```bash
cga create [OPTION]
```

| Option | Description                                              | Type   | Default | Required? |
|--------|----------------------------------------------------------|--------|---------|-----------|
| `-t`   | Enables to define custom backend and frontend templates. | `bool` | `false` | No        |

## 📝 Production-ready project templates

### Backend

- Backend template with Golang built-in [fe-service][net_http_url] package:
  - [`fe-service`][cga_net-http-template_url] — simple REST API with CRUD 
    and JWT auth.

## 🏆 A win-win cooperation

And now, I invite you to participate in this project! Let's work **together** to
create the **most useful** tool for developers on the web today.

Together, we can make this project **better** every day! 😘

## ⚠️ License

[`Create Go App CLI`][repo_url] is free and open-source software licensed under 
the [Apache 2.0 License][repo_license_url]. Official [logo][repo_logo_url] was 
created by [Kent Duong][author] and distributed under 
[Creative Commons][repo_cc_url] license (CC BY-SA 4.0 International).

<!-- Go -->

[go_download_url]: https://golang.org/dl/
[go_install_url]: https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies
[go_version_img]: https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go
[go_dev_url]: https://pkg.go.dev/github.com/ghn-toolkit/test-cli

<!-- Repository -->

[repo_url]: https://github.com/ghn-toolkit/cli
[repo_logo_url]: https://github.com/ghn-toolkit/cli/wiki/Logo
[repo_logo_img]: https://github.com/ghn-toolkit/cli/assets/11155743/95024afc-5e3b-4d6f-8c9c-5daaa51d080d
[repo_license_url]: https://github.com/ghn-toolkit/cli/blob/main/LICENSE
[repo_cc_url]: https://creativecommons.org/licenses/by-sa/4.0/
[repo_issues_url]: https://github.com/ghn-toolkit/cli/issues
[repo_pull_request_url]: https://github.com/ghn-toolkit/cli/pulls
[repo_discussions_url]: https://github.com/ghn-toolkit/cli/discussions
[repo_wiki_url]: https://github.com/ghn-toolkit/cli/wiki
[repo_wiki_faq_url]: https://github.com/ghn-toolkit/cli/wiki/FAQ


<!-- Author -->

[author]: https://github.com/comehere127
