#!/bin/sh

# Create .env file.
cp .env.example .env;

# Make compile script executable.
chmod 755 ./compile.sh;

cp config.default.yml config.yml;