##############################################################################
# BASIC TARGETS TO BUILD AND RUN THE PROGRAM
##############################################################################

# Standard env
GO ?= /usr/bin/go
MARKDOWN ?= /usr/bin/markdown
LIB_OPTIONS ?= -linkshared
BUILD_OPTIONS ?= -race

all: clean build

clean:
	rm -rf lib bin

# Install go standard librairies as shared
shared-install:
	@echo -n 'Installing standard librairies (as shared): '
	@sudo $(GO) install -buildmode=shared std && echo 'OK' || echo 'Fail'

build:
	$(GO) build $(BUILD_OPTIONS) -o bin/graal src/main.go

# For now, the 'home page' is generated using the README.md,
# until we'll use the source code to generate docs
docs:
	$(MARKDOWN) README.md >docs/home.html

run:
	./bin/graal --home=./docs/home.html


##############################################################################
# INSTALL NEEDED PACKAGES FOR DOCKER
##############################################################################
packages:
	@echo 'Installing needed packages (docker/debootstrap/etc.)'
	sudo apt-get -qq install debootstrap docker.io

##############################################################################
# CREATE THE FULL DOCKER IMAGES FOR TESTING
##############################################################################

dirs:
	@echo "Creating temporary/working directories"
	test -d tmp || mkdir tmp
	test -d instances || mkdir instances
	test -d tests || mkdir tests
	test -d tests/images || mkdir tests/images
	test -d tests/temp || mkdir tests/temp

# Dynamically generate a new SSH key used for deployment
ssh-key:
	rm -f deploy/id_rsa deploy/id_rsa.pub
	ssh-keygen -N '' -q -C 'graal-auth-key' -t rsa -b 2048 -f deploy/id_rsa

# Create a full Debian Jessie image, using debootstrap
# This is needed only once, to run the tests on a Debian machine
image-jessie: packages dirs ssh-keys
	@echo "Creating image: Debian Jessie, this may take a while, have a coffee…"
	sudo debootstrap jessie instances/jessie/
	sudo mkdir instances/jessie/root/.ssh
	sudo cp deploy/id_rsa.pub instances/jessie/root/.ssh/authorized_keys
	sudo bash -c "cd instances/jessie/ && tar cf ../../tmp/jessie.tar ."
	sudo chown $(USER):$(USER) tmp/jessie.tar
	mv tmp/jessie.tar tests/images/
	sudo rm -rf jessie

##############################################################################
# DOCKER TARGETS
##############################################################################

# Import the Jessie image
docker-jessie-import:
	sudo docker import - debootstrap/jessie < tests/images/jessie.tar

# Build jessie
docker-jessie-build:
	sudo docker build -t=graal-server - < deploy/debian-jessie.runit

# Run jessie
docker-jessie-start:
	sudo docker run -d -p 2222:22 -p 8123:80 --cidfile=tests/temp/graal-server.id graal-server

# Export the Jessie image into an image.
docker-jessie-export:
	sudo docker export `cat tests/temp/graal-server.id` > tests/images/graal-server.tar

# Commit changes on the docker
docker-jessie-commit:
	sudo docker commit `cat tests/temp/graal-server.id` graal-server

# stop jessie
docker-jessie-stop:
	sudo docker stop -t 30 `cat tests/temp/graal-server.id`
	sudo rm -f tests/temp/graal-server.id

# Connect as root on the vm
docker-jessie-connect:
	ssh -i deploy/id_rsa -p 2222 root@localhost

