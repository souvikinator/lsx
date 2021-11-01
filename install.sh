#!/usr/bin/env bash

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

cp "script/lsx.sh" "$HOME/.config/lsx/lsx.sh"

# build
go build -o "$GOPATH/bin/ls-x"

echo "INFO: build successful!"

cat <<-END

	All set, just add the following lines at the end of your shell resource file (.zshrc, .bashrc ....)

		source ~/.config/lsx/lsx.sh

	then restart your terminal.
----------------------------------------------	
	https://github.com/souvikinator/lsx
	If you liked the project, then feel free to drop a star :)
END
