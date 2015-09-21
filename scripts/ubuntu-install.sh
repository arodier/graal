#!/bin/sh

# Install required PHP packages
sudo apt-get install php-pear php5-dev
sudo echo no | pecl install mongo
sudo echo "extension=mongo.so" >/etc/php5/conf.d/mongo.ini

# Install rabbitmq server
sudo add-apt-repository "deb http://www.rabbitmq.com/debian/ testing main"
wget http://www.rabbitmq.com/rabbitmq-signing-key-public.asc
sudo apt-key add rabbitmq-signing-key-public.asc
sudo apt-get -q update
sudo apt-get -q -yy install rabbitmq-server
sudo rabbitmq-plugins enable rabbitmq_management
sudo service rabbitmq-server restart

exit 0
