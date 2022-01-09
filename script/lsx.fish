function lsx --description 'Run LSX'
    set LSX "$GOPATH/bin/ls-x"
    set TEMP_FILE "$HOME/.config/lsx/lsx.tmp"
    set LSX_BANNER "
    __    _  __
   / /___| |/ /
  / / ___/   / 
 / (__  )   |  
/_/____/_/|_|tended

github.com/souvikinator/lsx
"

    if test (count $argv) -lt 1
        $LSX
        cd (cat $TEMP_FILE)
        true >$TEMP_FILE
    else
        if test "$argv" = --help
            echo $LSX_BANNER
            $LSX --help
        else
            $LSX $argv
            cd (cat $TEMP_FILE)
            true >$TEMP_FILE
        end
    end
end
