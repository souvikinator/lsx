<h1 align="center">
  <br>
<a href="https://github.com/souvikinator/lsx"><img src="https://github.com/souvikinator/lsx/raw/master/assets/lsx-logo.svg" alt="lsx" width="400" height="350"></a>
<br>

</h1>

<h3 align="center">Navigate through terminal like a pro 😎 </h3>
<p align="center">
  <a href="https://opensource.org/licenses/">
    <img src="https://img.shields.io/badge/licence-MIT-brightgreen"
         alt="license">
  </a>
  <a href="https://github.com/souvikinator/lsx/issues"><img src="https://img.shields.io/github/issues/souvikinator/lsx"></a>
<a href="https://codebeat.co/projects/github-com-souvikinator-lsx-master"><img alt="codebeat badge" src="https://codebeat.co/badges/08315931-e796-4828-bfb0-18b6750d6f2a" /></a>
  <img src="https://img.shields.io/badge/made%20with-Go-blue">
  <img src="https://goreportcard.com/badge/github.com/souvikinator/lsx" alt="go report card" />
</p>

<p align="center">
	<a href="#-Demo">💻 Demo</a> •
  <a href="#%EF%B8%8F-install">⚗️ Installation</a> •
	<a href="#-contribution">🐜 Contribution</a> •
	<a href="#known-issues"> ❗Known Issues </a>
</p>

# ❓ Why?

It's a pain to `cd` and `ls` multiple times to reach desired directory in terminal. **ls-Xtended (lsx)** solves this problem by allowing users to smoothly navigate and search directories on the go with just one command. It also allows to create alias for paths making it easier for users to remember the path to the desired directory.

# 💻 Demo

> **Note**: once you reach the desired destination, use `ctr+c` to exit and stay in the desired destination

## Navigate through terminal and perform search:

- use `/` to trigger search and start typing to search

```bash
lsx
```

![lsx](https://github.com/souvikinator/lsx/blob/master/assets/demo.gif)

## Show hidden files as well

```bash
lsx -a
```

![lsx](https://github.com/souvikinator/lsx/blob/master/assets/all-mode.gif)

## Set **alias** for directory paths

```bash
lsx set-alias -n somealias -p path/to/be/aliased
```
or
```bash
lsx set-alias --path-name somealias --path path/to/be/aliased
```

![lsx](https://github.com/souvikinator/lsx/blob/master/assets/set-alias.gif)

## Updating Alias

`set-alias` can also be used to update any existing alias. Let's say alias `abc` already exists for path `a/b/c`. on can update it like so:

```bash
lsx set-alias -n abc -p d/e/f
```

## List **alias** created by user

```bash
lsx alias
```

![lsx](https://github.com/souvikinator/lsx/blob/master/assets/list-alias.gif)

## Remove existing **alias** 

```bash
lsx remove-alias aliasname
```

![lsx](https://github.com/souvikinator/lsx/blob/master/assets/remove-alias.gif)

# 📋 Todo

- [x] make a logo
- [ ] add icons
- [x] `-a`/`--all` mode
- [x] search
- [x] allow User can navigate to previous directory from the one they started ([#1](https://github.com/souvikinator/lsx/issues/))

# ⚗️ Install

> ## **Note**: make sure to have Go installed and your GOPATH is added to [PATH](https://stackoverflow.com/questions/21001387/how-do-i-set-the-gopath-environment-variable-on-ubuntu-what-file-must-i-edit)

### Step-1:

Clone the repo:

`git clone https://github.com/souvikinator/lsx.git`

### Step-2:

> `cd lsx`

> `chmod u+x install.sh`

> `./install.sh`

after installation is successful, add the following line at the end of your **current running shell resource file** (.zhsrc, .bashrc ...)

> `source ~/.config/lsx/lsx.sh`

and restart your terminal. Enjoy!

**Note**: Feel free to open an issue if any problems faced during installation.

If you liked the project, feel free to drop a star :)

# 🐜 Contribution

You can improve this project by contributing in following ways:

- report bugs
- fix issues
- request features
- asking questions (just open an issue)

and any other way if not mentioned here.

# ❗Known Issues

As of now the installation process is painful and the reason is a program runs as a child process in a terminal so eveything happens withing that child process. When we change the directory from go program the directory changes for that executable or to be specific "for that child process" and not of the shell. Which is why one needs to source a script in their shell resource file (`.zshrc`, `.bashrc`...).

The script contains a bash function as a wrapper around the lsx binary to make the whole `cd` thing work. This is what is prevent lsx to be distributed using some package manager. 

If anyone can comeup with something then feel free to open issue.
