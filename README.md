# Husky

[![Build Status](https://github.com/basti1302/go-husky/workflows/Go/badge.svg?branch=main)](https://github.com/basti1302/go-husky/actions?query=branch%3Amain)
[![Release](https://img.shields.io/github/release/basti1302/go-husky.svg)](https://github.com/basti1302/go-husky/releases)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/basti1302/go-husky)
[![Go Report Card](https://goreportcard.com/badge/github.com/basti1302/go-husky)](https://goreportcard.com/report/github.com/basti1302/go-husky)
![GitHub](https://img.shields.io/github/license/basti1302/go-husky)
![GitHub issues](https://img.shields.io/github/issues/basti1302/go-husky)

**This is a fork of <https://github.com/automation-co/husky>, adding support for git worktrees.**

Inspired by [husky.js](https://github.com/typicode/husky)

## Docs

### Installation

```
go install github.com/basti1302/go-husky@latest
```

### Getting Started

*Warning:* All `husky` commands are *potentially destructive*, they will *delete* all files in
`.git/hooks` and recreate them from `.husky/hooks`.
If you already have custom hooks in `.git/hooks`, make sure to create a backup of the hooks before
starting to use `husky`.

*Note:* All commands are supposed to be run from within your git repository.
If you use worktrees, in particular with a bare repository, run the commands from within a worktree, not the bare repository root.

Initialize husky by running `husky init`.

This will create the `.husky/hooks` folder with an example pre-commit hook.

You can add more hooks by running:

```bash
husky add <hook> "
  <your commands for that hook>
"
```

You can also create and edit files within `.husky/hooks`.

The `.husky/hooks` directory is supposed to be under version control, that is, you should commit and push it.

### Example

```bash
husky add pre-commit "
  go build -v ./... 
  go test -v ./...
"
```

If you have made any other changes in the hooks you can apply them by running `husky install`.

---

## Blogs and Resources
- [ Get Started with Husky for go ](https://dev.to/devnull03/get-started-with-husky-for-go-31pa)
- [ Git Hooks for your Golang project ](https://dev.to/aarushgoyal/git-hooks-for-your-golang-project-1168)

---

## Get Familiar with Git Hooks

Learn more about git hooks from these useful resources:
- [ Customizing Git - Git Hooks ](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks)
- [ Atlassian Blog on Git Hooks ](https://www.atlassian.com/git/tutorials/git-hooks)
- [ Fei's Blog | Get Started with Git Hooks ](https://medium.com/@f3igao/get-started-with-git-hooks-5a489725c639)

---

### Other Alternatives

If you feel husky does not fulfill your needs you can also check out:
- https://github.com/typicode/husky
- https://pre-commit.com/

---

<div align="center">

Developed originally by [@automation-co](https://github.com/automation-co), with additions by
[@basti1302](https://github.com/basti1302).
