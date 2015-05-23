#!/bin/bash

exec >> /tmp/snotify.log 2>&1
echo ""
date
echo "$0 $*"

id="snotify_popup"

if [ "$1" == "--kill" ]; then
    pkill -f ${id}
    exit 0
fi

# TODO: handle more num args errors
[ $# -gt 1 ] && {
    title=$1
    shift
}

lineheight=21
numlines=$(echo "$@" | xargs -n3 | wc -l)
(( height = (${numlines} * ${lineheight}) + 10 ))

(
echo " ${title}"
echo " $@" | xargs -n3 | sed 's/.*/ & /'
) | dzen2 -p -title-name ${id} \
                     -fn '-*-terminus-bold-r-*-*-14-*-*-*-*-*-*-*' \
                     -fg "#c0c0c0" -bg "#1b1d1e" \
                     -l ${numlines} -w 210 -y -${height} -x -250 \
                     -sa 'r' -ta 'c' \
                     -e 'raise;onnewinput=uncollapse;button1=exit;sigusr1=exit'&
                     # -e 'onstart=hide;onnewinput=uncollapse;button1=exit;sigusr1=exit'&
sleep .025s && transset-df --name ${id} .7 >/dev/null 2>&1
