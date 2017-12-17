# What is this?

Syncer is a project that makes it easy to run rsync for multiple source/destination
directory pairs at the same time.

# How to install

## Prerequisites
Make sure that you have [Golang](https://golang.org) and rsync installed
on your machine.

## Steps

1. Clone the repository to your $GOPATH:
    `git clone https://github.com/sirrah23/Syncer.git`
2. Run `go install Syncer/Syncer` at the root of your $GOPATH
3. Navigate to the bin directory at your $GOPATH to find the Syncer executable

## Usage

Say that you want to rsync the following directories:

    /home/user/DIR1 -> /home/user/backup/DIR1
    /home/user/DIR2 -> /home/user/backup/DIR2
    /home/user/DIR3 -> /home/user/backup/DIR3

Create a file called ''files.csv'' with the content:

    /home/user/DIR1,/home/user/backup/DIR1
    /home/user/DIR2,/home/user/backup/DIR2
    /home/user/DIR3,/home/user/backup/DIR3

and then run `Syncer -files=files.csv`.


After you do that all of the source directories will have their
contents rsync'ed to all of the destination directories.
