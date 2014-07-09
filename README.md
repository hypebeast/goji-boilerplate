# !!Work in Progress!!
**More coming soon!**

# goji-boilerplate

goji-boilerpalate is a project skeleton for Goji web framework, Twitter Bootstrap SASS, Font Awesome and Gulp. It also promotes seperation of concerns and app re-usability. This makes building middle and big size projects easier.

## Features

General features:

  * Goji web framework
  * Separated templates for base layout, header, navbar and footer. 
  * Bower for client side package management
  * Gulp for compiling SASS to CSS and for minification and uglifying the Javascript code

Client side frameworks:

  * Twitter Bootstrap (SASS version)
  * Font Awesome
  * Gulp

## Getting Started

### Prerequisites

  * [Go](http://golang.org/)
  * [NodeJS + NPM](http://nodejs.org/)
  * Make sure that $GOPATH is set

### Installation

The easiest way to get started is to *go get* the repository:

```bash
# Get the latest version from Github (make sure *$GOPATH* is set)
go get github.com/hypebeast/goji-boilerplate.git

cd $GOPATH/src/github.com/hypebeast/goji-boilerplate

# Install NPM dependencies
npm install

# Install all Bower components
bower install

# Run Gulp
gulp

# Build the app
go build

# Run the app
./goji-boilerplate
```

## Directory Structure

TODO

## Future Development

  * Registration/Login Page
  * Add environment depending configuration
  * Add support for database and models
  * Add support for sessions
  * Live reload utility for Go (e.g. gin)
  * Panic recovery
  * Logger
  * Client-Side MVC Framework: AngularJS?

## Contribute

  * Fork [goji-boilerplate](https://github.com/hypebeast/goji-boilerplate)
  * Make some changes
  * Send your Pull Request
  * Thanks :)

## Credits

-

## License

The MIT License (MIT)

Copyright (c) 2014 Sebastian Ruml

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

