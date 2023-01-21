#!/usr/bin/env bash

# set the repo URL
repo="https://github.com/donuts-are-good/fir-server.git"

# set the project directory
dir="$HOME/fir/fir-server"

# set the binary name
binary="fir-server"

# check if it already exists
if [ ! -d "$dir" ]; then

  # clone if it doesn't exist
  git clone "$repo" "$dir"

# end the loop
fi

# now that it should be there, enter the project directory
cd "$dir"

# check for changes
while true; do
  
  # fetch changes
  git fetch

  # check for new changes
  if git rev-parse HEAD..origin/master --quiet; then
    
    # pull them in
    git pull

    # check if any project files have been modified
    if [ -n "$(git diff --name-only HEAD..HEAD@{1})" ]; then

      # compile the new copy
      go build 

      # wait 2 seconds just in case
      sleep 2

      # kill the current running copy
      pkill "$binary"
      
    fi
  fi

  # cycle time 10s
  sleep 10

done
