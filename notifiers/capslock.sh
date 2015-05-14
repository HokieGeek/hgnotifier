#!/bin/sh

case $1 in
    --show)
        echo "A" | dzen2 -p -title-name 'osdnotifier_caps' \
                            -fn '-*-terminus-bold-r-*-*-15-*-*-*-*-*-*-*' \
                            -fg "#005f00" -bg "#000000" \
                            -w 20 -y -30 -x -35 \
                             -e 'raise'&
        sleep .1s && transset-df --name 'osdnotifier_caps' .65 >/dev/null 2>&1
        ;;
    --hide) pkill -f osdnotifier_caps ;;
esac
