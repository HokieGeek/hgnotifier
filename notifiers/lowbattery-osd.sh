#!/bin/bash

id="hgnotifier_lowbattery"

function createConkyFile() {
cat << EOF >> $1
background yes
out_to_console yes
out_to_x no
# Update interval in seconds
update_interval 1

TEXT
\${texeci 1 date +%H:%M:%S }
# \${texeci 1 acpi | head -1 | sed 's/.* \([0-9:]*\) .*/\1/g' }
EOF
}

conkyfile=/tmp/`basename $0`.conky
case $1 in
    on) [ ! -f "${conkyfile}" ] && createConkyFile ${conkyfile}
        conky -b -c ${conkyfile} | dzen2 -p -title-name ${id} \
                                         -fn '-*-terminus-bold-r-*-*-15-*-*-*-*-*-*-*' \
                                         -fg "#ff0000" -bg "#1b1d1e" \
                                         -w 75 -y 14 -x -70 \
                                         -e 'raise'&
        sleep .025s && transset-df --name ${id} .65 >/dev/null 2>&1
        ;;
    off) pkill -f `basename ${conkyfile}`
         pkill -f ${id}
         rm -rf ${conkyfile}
         ;;
esac
