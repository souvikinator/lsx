#!/usr/bin/env bash

ZSH_DIR="$HOME/.oh-my-zsh"
ZSH_FUNC_DIR="$ZSH_DIR/functions"

FISH_DIR="$HOME/.config/fish"
FISH_FUNC_DIR="$FISH_DIR/functions"

# git installed?
if ! [[ -x "$(command -v git)" ]];then
	echo "ERROR: Git not installed!"
	exit 1
fi

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
if [[ -d "$HOME/.config/fish/functions" ]];then
  cp "script/lsx.fish" "$HOME/.config/fish/functions/lsx.fish"
fi


#zsh
if [[ -d "$ZSH_DIR" ]];then
	mkdir -p "$ZSH_FUNC_DIR"
  cp "script/lsx.sh" "$ZSH_FUNC_DIR/lsx"
fi

#fish
if [[ -d "$FISH_FUNC_DIR" ]];then
  cp "script/lsx.fish" "$FISH_FUNC_DIR/lsx.fish"
fi

# build
go build -o "$GOPATH/bin/ls-x"

echo "INFO: build Success!"

cat <<-END

ZSH users, add following line at the end of .zshrc 

		autoload -Uz lsx

BASH users, add following line at the end of .bashrc 

		source ~/.config/lsx/lsx.sh

and restart the shell

----------------------------------------------	

https://github.com/souvikinator/lsx
If you liked the project, then feel free to drop a star :)
END
