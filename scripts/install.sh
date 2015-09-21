#!/bin/sh

# Install required packages
required='make rabbitmq-server php5-json'

# Install the required packages
sudo apt-get -q update
sudo apt-get -yy install $required

exit 0
