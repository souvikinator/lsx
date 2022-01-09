#!/usr/bin/env bash

BASHRC="$HOME/.bashrc"

ZSHRC="$HOME/.zshrc"
ZSH_DIR="$HOME/.oh-my-zsh"
ZSH_FUNC_DIR="$ZSH_DIR/functions"

FISH_DIR="$HOME/.config/fish"
FISH_FUNC_DIR="$FISH_DIR/functions"

FISH_SHELL_SCRIPT_URL="https://raw.githubusercontent.com/souvikinator/lsx/master/script/lsx.fish";
BASH_SHELL_SCRIPT_URL="https://raw.githubusercontent.com/souvikinator/lsx/master/script/lsx.sh";

# go installed?
if ! [[ -x "$(command -v go)" ]];then
	echo "ERROR: GO not installed!"
	exit 1
fi

#check if gopath is set
if [[ -z "$GOPATH" ]];then
	echo "ERROR: Your GOPATH is not set."
	exit 1
fi

if ! [[ -d "$HOME/.config/lsx" ]];then
	mkdir -p "$HOME/.config/lsx"
fi

echo "Downloading required scripts...";


if [[ -d "$ZSH_DIR" ]];then
	#zsh
	mkdir -p "$ZSH_FUNC_DIR"
  curl "$BASH_SHELL_SCRIPT_URL" -o "$ZSH_FUNC_DIR/lsx"
elif [[ -d "$FISH_FUNC_DIR" ]];then
	#fish
  curl "$FISH_SHELL_SCRIPT_URL" -o "$FISH_FUNC_DIR/lsx.fish"
else
	# bash
	curl "$BASH_SHELL_SCRIPT_URL" -o "$HOME/.config/lsx/lsx.sh"
fi


# install
go install "github.com/souvikinator/lsx@latest"


if [[ $? -ne 0 ]];then
  echo "ERROR: Installation failed!"
  exit 1
else
  echo "INFO: Successfully installed!"
fi

# change installed binary name: lsx -> ls-x
echo "changing executable name"
cp "$GOPATH/bin/lsx" "$GOPATH/bin/ls-x"

if [[ $? -ne 0 ]];then
  echo "ERROR: failed to change executable name!"
  exit 1
else
  echo "INFO: executable name changed!"
fi

# add line to bashrc if exists
if [[ -f "$BASHRC" ]];then
	if ! [[ $(grep -e "source ~/.config/lsx/lsx.sh" "$BASHRC") ]];then
		echo "INFO: Adding 'source ~/.config/lsx/lsx.sh' in $BASHRC"
		echo "#lsx" >> "$BASHRC"
		echo "source ~/.config/lsx/lsx.sh" >> "$BASHRC"
	fi
fi

# add line to zshrc if exists
if [[ -f "$ZSHRC" ]];then
	if ! [[ $(grep -e "autoload -Uz lsx" "$ZSHRC") ]];then
		echo "INFO: Adding 'autoload -Uz lsx' in $ZSHRC"
		echo "#lsx" >> "$ZSHRC"
		echo "autoload -Uz lsx" >> "$ZSHRC"
	fi
fi

echo "All set! Restart your terminal."
