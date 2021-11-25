lsx () {
	TEMP_FILE="$HOME/.config/lsx/lsx.tmp"

	LSX_BANNER="
		 __    _  __
		/ /___| |/ /
	 / / ___/   / 
	/ (__  )   |  
 /_/____/_/|_|tended

	github.com/souvikinator/lsx
	"

	ARGS=("$@")
	ARGS_LEN="${#ARGS[@]}"
	
	if [[ ARGS_LEN -lt 1 ]];then
		ls-x
		LSX_CWD=$(cat "$TEMP_FILE")
		cd "$LSX_CWD"
		true > "$TEMP_FILE"
	else
		if [[ "${ARGS[0]}" == "--help" ]];then
			echo "$LSX_BANNER"
			ls-x --help
			return
		else
			ls-x "${ARGS[@]}"
			LSX_CWD=$(cat "$TEMP_FILE")
			cd "$LSX_CWD"
			true > "$TEMP_FILE"
		fi
	fi
}
