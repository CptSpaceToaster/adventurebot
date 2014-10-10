Installation Notes
========

Go (usually referred to as 'golang' on the internet for all intensive porpoises) can (and should) manage dependencies and the build process, so while it seems just like another language with compiler, it actually can get you from project->executable in a couple of steps

Installing the Go Toolset
========
You will be prefacing most of everthing you accomplish in go by invoking `go <command> <arguments>` (big surprise right?!?)

To do that, you'll need to install go tools (at least version 1.3.1).
Currently, apt-get hands 'buntu users an older version, I'm not sure about the state of other package managers, but manually installing isn't bad

Pull a tarball from [Go's download page](http://golang.org/dl/) and follow their [installation instructions](http://golang.org/doc/install) up until you can type `go version` or `go help`

`export VERSION='1.3.3'`  
`export OS='linux'`  
`export ARCH='amd64'`  

`wget golang.org/dl/go$VERSION.$OS-$ARCH.tar.gz`  
`sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz`  
`export PATH=$PATH:/usr/local/go/bin` - You probably want this in your .profile  
Make sure to create a directory to store your go-projects, **and** set the `$GOPATH` to the location of the the directory you create  
`mkdir ~/go-proj-dir`  
`export GOPATH=~/go-proj-dir` - Probably want this in your .profile as well  
Check to see if the tool exists!  
`go version`  
`> go version go1.3.3 linux/amd64`

If you're interested diving into go, you may want to follow further steps and create everyone's favorite "hello world" program.  I'll leave the initial tinkering as an excersize for those who care.

More information: `go help` especially `go help gopath`

Now you're ready to gather dependencies and the projects you want to work with.

Obtaining Go Depdencies and Projects
========
Go wrapped the acquisition of projects and repositories using the command `go get`.  At first I wanted to complain, because I wanted to `git clone` my projects into their place, and manage where I was storing code in my own folder structure... then I realized why the creators did this:

1. `go get` can pull from various repositories... whether you use mercurial, git, SVN, etc, you can usually just point to `go get <address>` and it will pull **a working copy** of the repository for you (and build it automagically).

2. Dependencies can be obtained with the same command, meaning that developers can easily rely on projects that are maintained in an entirely different version control system.

Complain if you want to... but this should mean all projects can rely on developers having the same folder structure within their `$GOPATH`.  Dependencies just work...

For this project, you'd need both:  
`go get github.com/gorilla/schema`  
`go get github.com/cptspacetoaster/adventurebot`  

Don't download them in reverse order...

You should be able to find this repo in
`cd $GOPATH/src/github.com/cptspacetoaster/adventurebot`

Build Stuff!
========
Guess what?  When you ran `go get` you should have automagically built the executable that the project generates!

Did you make changes to the source?  Rebuild a project using `go install <project>` (Assuming the project you pulled was stable)  
`go install github.com/cptspacetoaster/adventurebot`  

Your executable(s) can be found here:  
`cd $GOPATH/bin`  

Now head back to the mainpage to finish reading that [README](https://github.com/CptSpaceToaster/adventurebot)... you're only a configuration file away!
