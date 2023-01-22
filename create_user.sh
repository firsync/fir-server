#!/usr/bin/env bash

# get the public key string
pubkey=$2

# remove any trailing characters after the 64th character
pubkey=`echo $pubkey | sed 's/\(^.\{64\}\).*/\1/'`

# create a new user with the public key as the username
useradd $pubkey

# create a home directory for the new user
homeDir="/home/$pubkey"
mkdir $homeDir

# set ownership of the home directory to the new user
chown $pubkey:$pubkey $homeDir

# create an authorized_keys file for the new user
ssh-keygen -f "$homeDir/.ssh/authorized_keys" -t ed25519 -N ""

# add the public key to the authorized_keys file
echo $2 >> "$homeDir/.ssh/authorized_keys"

# set the correct permissions on the authorized_keys file
chmod 600 "$homeDir/.ssh/authorized_keys"
