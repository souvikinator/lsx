#!/usr/bin/env bash

BASHRC="$HOME/.bashrc"

ZSHRC="$HOME/.zshrc"
ZSH_DIR="$HOME/.oh-my-zsh"
ZSH_FUNC_DIR="$ZSH_DIR/functions"

FISH_DIR="$HOME/.config/fish"
FISH_FUNC_DIR="$FISH_DIR/functions"

# go installed?
if ! [[ -x "$(command -v go)" ]];then
	echo "ERROR: GO not installed!"
	exit 1
fi

if ! [[ -d "$HOME/.config/lsx" ]];then
	mkdir -p "$HOME/.config/lsx"
fi

# bash
cp "script/lsx.sh" "$HOME/.config/lsx/lsx.sh"

#zsh
if [[ -d "$ZSH_DIR" ]];then
	mkdir -p "$ZSH_FUNC_DIR"
  cp "script/lsx.sh" "$ZSH_FUNC_DIR/lsx"
fi

#fish
if [[ -d "$FISH_FUNC_DIR" ]];then
  cp "script/lsx.fish" "$FISH_FUNC_DIR/lsx.fish"
fi

#check if gopath is set
if [[ -z "$GOPATH" ]];then
	echo "ERROR: Your GOPATH is not set."
	exit 1
fi


# build
go build -o "$GOPATH/bin/ls-x"

if [[ $? -ne 0 ]];then
  echo "ERROR: build Failure!"
  exit 1
else
  echo "INFO: build Success!"
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
