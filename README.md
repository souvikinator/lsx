<h1 align="center">
  <br>
<a href="https://github.com/souvikinator/lsx"><img src="https://github.com/souvikinator/lsx/raw/master/assets/lsx-logo.svg" alt="lsx" width="600"></a>
<br>

</h1>

<h3 align="center">Navigate through terminal like a pro!</h3>
<p align="center">
  <a href="https://opensource.org/licenses/">
    <img src="https://img.shields.io/badge/licence-MIT-brightgreen"
         alt="license">
  </a>
  <a href="https://github.com/souvikinator/lsx/issues"><img src="https://img.shields.io/github/issues/souvikinator/lsx"></a>
  <img src="https://img.shields.io/badge/made%20with-Go-blue">
  <img src="https://goreportcard.com/badge/github.com/souvikinator/lsx" alt="go report card" />
</p>

<p align="center">
	<a href="#-Demo">💻 Demo</a> •
  <a href="#%EF%B8%8F-install">⚗️ Installation</a>
</p>

# 💻 Demo

In simple words **lsx** or **ls xtended** is combination of two commands: `ls` and `cd`. It lets you navigate through terminal with ease, somewhat similar to file explorer.

> **Note**: once you reach the desired destination, use `ctr+c` to exit and stay in the desired destination

![lsx](https://github.com/souvikinator/lsx/blob/master/assets/demo.gif)

# 📋 Todo

- [ ] make a logo
- [ ] add icons
- [x] `-a`/`--all` mode
- [x] search
- [x] allow users to go back (user cannot go to the previous directory of the directory they started from, ref next todo)
- [ ] allow User can navigate to previous directory from the one they started ([#1](https://github.com/souvikinator/lsx/issues/))
- [ ] Display files as well and allow various operations on them (renaming, bulk renaming)

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

If you liked the project, feel free to drop a star :)
