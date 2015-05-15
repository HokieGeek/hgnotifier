#!/bin/sh

id="snotify_capslock"

case $1 in
    on)
                            # -fg "#005f00" -bg "#000000" \
        echo "A" | dzen2 -p -title-name ${id} \
                            -fn '-*-terminus-bold-r-*-*-15-*-*-*-*-*-*-*' \
                            -fg "#ffffff" -bg "#ff0000" \
                            -w 20 -y -35 -x -38 \
                             -e 'raise'&
        sleep .025s && transset-df --name ${id} .65 >/dev/null 2>&1
        ;;
    off) pkill -f ${id} ;;
esac
