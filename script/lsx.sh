#!/bin/bash
# TODO: work on arguments
# --help, --version
lsx () {
	ls-x
	LSX_CWD=$(cat "$HOME/.lsx.tmp")
	echo "$LSX_CWD"
	cd "$LSX_CWD"
}
