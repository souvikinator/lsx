read -r -d '' LSX_BANNER <<- EOM
	_
 | |_____  __
 | / __\ \/ /
 | \__ \>  < 
 |_|___/_/\_\\

 https://github.com/souvikinator/lsx
EOM

lsx () {
	OPTION="$1"
	if [[ "$OPTION" == "--help" || "$OPTION" == "-h" ]]; then
		printf "\n%s \n\n" "$LSX_BANNER"
		ls-x --help
		return
	elif [[ "$OPTION" == "--version" || "$OPTION" == "-v" ]]; then
		ls-x --version
		return
	else
		ls-x "$OPTION"
		LSX_CWD=$(cat "$HOME/.lsx.tmp")
		cd "$LSX_CWD"
	fi
}