TEMP_FILE="$HOME/.config/lsx/lsx.tmp"

lsx () {
	ARGS=("$@")
	ARGS_LEN="${#ARGS[@]}"


	if [[ ARGS_LEN -lt 1 ]];then
		ls-x
		LSX_CWD=$(cat "$TEMP_FILE")
		cd "$LSX_CWD"
		true > "$TEMP_FILE"
	else
		ls-x $ARGS
		LSX_CWD=$(cat "$TEMP_FILE")
		cd "$LSX_CWD"
		true > "$TEMP_FILE"
	fi
}