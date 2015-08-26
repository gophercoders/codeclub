## GopherCoders.com Code Club Programs

These are the sample programs used in the GopherCoders.com code clubs.
The order of teaching is:

* helloclub
* hellofriends
* namesandages
* hellostrangers
* worldtemperature
* timesquestion
* timesquiz
* mandelbrot

Each program builds on the last and introduces new concepts.

* helloclub - introduces go, a go source code file and the `go` command.
* hellofriends - introduces string variables and output.
* namesandages - introduces number variables and maths go.
* hellostranger - introduces simple keyboard input.
* worldtemperature - introduces selection or `if` statements.
* timesquestion - introduces `if else` statements.
* timesquiz - introduces repetition or `for` statements.
* mandelbrot - brings everything together to plot a mandelbrot set.

Any program named `*-complete.go` contains a sample solution. The `*.go` programs
the pupils are expected to complete.

### Dependencies

The programs hellostranger, worldtemperature, timesquiz and timesquestion programs
depend on the `simpleio` package. The package is go get'able from
````
go get -u github.com/gophercoders/simpleio
````

Additionally the timesquiz and timesquestion programs also depend upon the `random`
package. The package is go get'able from

````
go get -u github.com/gophercoders/random
````

The mandelbrot package relies on the `go-sdl2` package. See [here](https://github.com/veandco/go-sdl2)
for the installation instructions.

Please see the [COPYRIGHT](https://github.com/gophercoders/codeclub/blob/master/COPYRIGHT)
file for copyright information.

Please see the [LICENSE](https://github.com/gophercoders/codeclub/blob/master/LICENSE)
file for license information.
