read -r -d '' BANNER << EOM
	_
 | |_____  __
 | / __\ \/ /
 | \__ \>  < 
 |_|___/_/\_\\
EOM

lsx () {
	OPTION="$1"
	if [[ "$OPTION" == "--help" || "$OPTION" == "-h" ]]; then
		printf "\n%s \n\n" "$BANNER"
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