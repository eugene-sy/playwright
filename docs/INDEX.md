# Playwright

[![Build Status](https://travis-ci.org/eugene-sy/playwright.svg?branch=master)](https://travis-ci.org/eugene-sy/playwright)
[![Go Report Card](https://goreportcard.com/badge/github.com/eugene-sy/playwright)](https://goreportcard.com/report/github.com/eugene-sy/playwright)

Simple utility that helps you manage Ansible roles folder structure.

## Why

Ansible is shipped with a [tool](http://docs.ansible.com/ansible/galaxy.html#create-roles) for creation of folder structure for new roles.
It creates full structure of folders and `main.yml` files. Generally you don't need to create all that folders and most frequent case is having `tasks` folder and `main.yml` there. Also, unused folders and files produce same effect as dead code in any codebase.

Playwright allows you to control which parts of the role you create and when.

## How to use it

You need to set up path to your `ansible.cfg` file.
By default playwright will expect to find it in:
- `ANSIBLE_CONFIG` environment variable
- in the current directory by name `ansible.cfg` or `.ansible.cfg`
- in your system config folder `/etc/ansible/ansible.cfg`

You need to set up path to the roles folder in your `ansible.cfg`:

```
roles_path=/somewhere/in/my/system
```

Now you can call playwright to build folder structure:

```
playwright [<flags>] <command> [<args> ...]

Flags:
  --help       Show context-sensitive help (also try --help-long and --help-man).
  --handlers   Add 'handlers' folder
  --templates  Add 'templates' folder
  --files      Add 'files' folder
  --vars       Add 'vars' folder
  --defaults   Add 'defaults' folder
  --meta       Add 'meta' folder
  --all        Apply action to all folders
  --version    Show application version.

Commands:
  help [<command>...]
	Show help.

  create <name>
	Creates a playbook

  update <name>
	Updates a playbook

  delete <name>
	Deletes a playbook
```

By default playwright creates only `tasks` folder and `main.yml` in it.

## Building and installing

To build `playwright` you need

- GoLang installed and `$GOPATH` set
- `glide` to install dependencies in future

To build and install run next command:

```
sudo make install
```

Binary file will be copied to your `/usr/local/bin` directory.

To simply build binary, run:

```
make build
```

## Changelog

Check `CHANGELOG.md` file.

## License

This software is built and distributed under GPLv3 license (Ansible uses this license).
For more information check `LICENSE.md` file in the root folder of the repository.

<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','https://www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-58752539-1', 'auto');
  ga('send', 'pageview');
</script>
