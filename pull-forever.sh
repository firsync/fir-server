#!/usr/bin/env bash

export TZ=America/Chicago

freshsincenow="`date +'%A, %B %d, %Y - %r'`"
gitlogstatus="`cd ~/fir/fir-server/ && git log | head -n 5 | tail -n 1`"
thislogo="
   ____                    
  / _(__________ _____ ____
 / _/ / __(_-/ // / _ / __/
/_//_/_/ /___\_, /_//_\__/ 
            /___/          

 \e[38;5;214mFir\e[38;5;208m Repository Monitor
 \e[38;5;7m
 \e[38;5;77mFresh since: \n\e[38;5;66m    "`date +'%A, %B %d, %Y - %r'`" \e[38;5;15m"

function thisthing {
  cd ~/webapp
  clear
  echo -e "\e[38;5;202m$thislogo" 
  echo -e "\n \e[38;5;77mTime now:\e[38;5;66m \n    "`date +'%A, %B %d, %Y - %r'`
  sleep 1 
}


function thatthing {
  cd ~/webapp
  clear
  echo -e "\e[38;5;202m$thislogo" 
  echo -e "\n \e[38;5;77mTime now:\e[38;5;66m \n    "`date +'%A, %B %d, %Y - %r'`
  sleep 1 
}


function theotherthing {
  cd ~/webapp
  clear
  echo -e "\e[38;5;202m$thislogo" 
  echo -e "\n \e[38;5;77mTime now:\e[38;5;66m \n    "`date +'%A, %B %d, %Y - %r'`
  if [[ $(git pull|grep 'Already up to date.') ]]
  then
    echo ""
  else 
    echo -e "\e[38;5;41mUpdating..." 
    killall fir-server
    thislogo="
   ____                    
  / _(__________ _____ ____
 / _/ / __(_-/ // / _ / __/
/_//_/_/ /___\_, /_//_\__/ 
            /___/          

 \e[38;5;214mFir\e[38;5;208m Repository Monitor
 \e[38;5;7m
 \e[38;5;77mFresh since: \n\e[38;5;66m    "`date +'%A, %B %d, %Y - %r'`" \e[38;5;15m"
 cp ~/fir/fir-server/pull-forever.sh ~/pull-forever.sh
fi 
}


while true ; do thisthing && thatthing && theotherthing ; done 