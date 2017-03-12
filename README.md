# Playwright

Simple utility that creates Ansible playbook folder structure for you.

## How to use it

1. You need to set up path to your `ansible.cfg` file.
By default playwright will expect to find it in:
- `ANSIBLE_CONFIG` environemnt variable
- in the current directory by name `ansible.cfg` or `.ansible.cfg`
- in your system config folder `/etc/ansible/ansible.cfg`

2. You need to set up path to the roles folder in your `ansible.cfg`:

```
roles_path=/somewhere/in/my/system
```

3. Now you can call playwright to build folder structure:

```
playwright my_awesome_playbook --with-handlers --with-templates --with-files --with-vars --with-defaults --with-meta
```

By default playwright creates only `tasks` folder and `main.yml` in it.

Flags starting with `--with` add named folders to the final structure

- `--with-handlers` - handlers folder and `main.yml` in it
- `--with-templates` - empty templates folder
- `--with-files` - empty files folder
- `--with-vars` - vars folder and `main.yml` in it
- `--with-defaults` - defaults folder and `main.yml` in it
- `--with-meta` - meta folder and `main.yml` in it

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

## License

This software is built and distributed under GPLv3 license (Ansible uses this license).
For more information check `LICENSE.md` file in the root folder of the repository.
