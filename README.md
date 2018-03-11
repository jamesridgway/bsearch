# bsearch
[![Build Status](https://travis-ci.org/jamesridgway/bsearch.svg?branch=master)](https://travis-ci.org/jamesridgway/bsearch)
[ ![Download](https://api.bintray.com/packages/jamesridgway/debian/bsearch/images/download.svg) ](https://bintray.com/jamesridgway/debian/bsearch/_latestVersion)

A utility for binary searching a sorted file for lines that start with the search key.

    NAME:
       bsearch - utility for binary searching a sorted file for lines that start with the search key

    USAGE:
       bsearch [options] SEARCH_KEY FILENAME

    VERSION:
       1.0.0

    COMMANDS:
         help, h  Shows a list of commands or help for one command

    GLOBAL OPTIONS:
       -r, --reverse      the reverse flag indicates the file is sorted in descending order
       -i, --ignore-case  case insensitive
       -t, --trim         ignore whitespace
       -n, --numeric      use numeric comparison
       --help, -h         show help
       --version, -v      print the version

## Installing
You can install bsearch via the following OS specific repositories

### Debian/Ubuntu

    # Add Bintray's GPG key
    sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 379CE192D401AB61

    # Add the repository
    echo "deb https://dl.bintray.com/jamesridgway/debian xenial main" | sudo tee -a /etc/apt/sources.list

    # Update apt
    sudo apt-get update

    # Install
    sudo apt-get install bsearch
