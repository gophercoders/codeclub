// The program plot a mandelbrot the famous set
package main

import (
	"fmt"
	"time"

	// the colours are defined in the file caled palette.go in the toolbox directory
	"github.com/gophercoders/toolbox/palette"
	// load the toolbox
	"github.com/gophercoders/toolbox"
)

// These are the variables for the graphics library
// They have to be outside of the main function because the functions at the
// end of the file need them.
// This the window we are going to draw into
var window toolbox.Window

// These constants are important. They are the width and height of the window
// A constant is just means you can't change these values once the
// program starts running
// If you change these you will change the size of the image
// try starting with 800 for the width and 600 for the height, or 640 by 480
const windowWidth = 1024
const windowHeight = 768

// plot stores the colour of every point on the screen
// Technically this is an "array" but you can just think of
// this as a very very long list!
// At each elment of in the array we store the colour of that point.
var plot [windowWidth][windowHeight]toolbox.Colour

// The programs main function
func main() {
	// ---- This is the start of the graphics setup code ----
	toolbox.Initialise()
	// defer is a go keyword and a special feature.
	// This means that go will automatically call the function toolbox.Close() before
	// the program exits for us. We don't have to remember to put this at the end!
	defer toolbox.Close()

	// Now we have to create the window we want to use.
	// We can use teh toolbox packae for this,
	window = toolbox.CreateWindow("Mandelbrot", windowWidth, windowHeight)
	// defer is a new keyword.
	// It just means that the DestroyWindow function will be called
	// automatically before the program exists. This ensures
	// that the window is correctly destroyed.
	defer toolbox.DestroyWindow(window)
	// ---- This is the end of the graphics setup code ----

	// These are the maximum and minimum values on the mandelbrot graph
	// float64 is the type go uses for decimal fractions
	var minRe float64 // this is the minumin in the x axis
	var maxRe float64 // this is the maximum in the x axis
	var minIm float64 // this is the minimum in the y axis
	var maxIm float64 // this is the maximum in the y axis - we calculate this

	// set the values for the minimum, and maximum in the x axis and the minimum
	// in the y axis
	minRe = -2.2
	maxRe = 1.0
	minIm = -1.2

	// calculate maxIm the maximum in the y axis. If we calculate this value we
	// can use the aspect ratio of the window to make sure the image is not
	// squashed or stretched.
	var interval float64
	interval = (maxRe - minRe)                                 // this is the length of the x axis
	var aspectRatio float64                                    // this is the aspect ratio of the window
	aspectRatio = float64(windowHeight) / float64(windowWidth) // make sure we do float64 division not int division. Otherwise we loose the decimal places!
	// now calcualte what maxIm should be
	maxIm = minIm + interval*aspectRatio

	// The scaling factors
	// The scaling in the x axis
	var reFactor float64
	// The scaling in the y axis
	var imFactor float64

	// calculate the scaling values
	reFactor = (maxRe - minRe) / float64(windowWidth-1)
	imFactor = (maxIm - minIm) / float64(windowHeight-1)

	// These are the screen (pixel) coordinates
	var x int
	var y int

	// this is the value of c - the value of the current pixel
	var cRe float64 // the real part
	var cIm float64 // the imaginary part

	// These are the loop variabels that control how many times we have done the
	// calculation
	// This is the number of times we have been around the calculation loop
	var loopCount = 0
	// This is the maximum number of times around the loop. This number needs to
	// be big. The bigger the number the more detail you see (up to the screen resolution)
	// but the longer the images takes to draw. Try values of 17, 35, 69
	var maxLoopCount = 138 // this is the size of the palette. Multiples of this number are good

	// this is the value of Z
	var zRe float64 // the real part
	var zIm float64 // the imaginary part

	// these are temporary variables which the algorithm uses to avoid
	// recalculating them
	var zReSqrd float64 // zRe * zRe
	var zImSqrd float64 // zIm * zIM

	// this is the colour of the pixel to plot
	// The new type comes form the toolbox. It represents colour as 4 numbers.
	// One for red, one for green, one for blue and one for transparancy.
	var colour toolbox.Colour

	// timing
	// The program time how long it takes to plot the mandelbrot set.
	// time.Time is a new type, that deals with timing. It measures time
	// in milliseconds - 1 millionth of a second
	var startTime time.Time
	startTime = time.Now()

	// The equation for the mandelbrot set is
	//
	//	next Z = (Z * Z) + c
	//
	//  c is the value of the point in the image/graph
	//
	// 	If we do this calculation HUNDREDS of times then if the absolute value of Z
	// 	is greater than 4 then c is not in the mandelbrot set. This gives us a coloured pixel.
	//	Otherwise c is in the mandelbrot set. This gives us a black pixel.
	//
	// The Algorithm is
	//
	// for every pixel in the image - working top to bottom, left to right
	// 		work out the value of c that the pixel represents
	//		copy the current values to c to z
	//		loop count = 0
	//		for loop count < MaxLoopCount
	//			if absolute value of Z > 4 		// this is the escape test!
	//				stop the loop
	//			work out the next Z value
	//			add one to loop count
	//		when the loop stops work out the colour - based on the number of times around the loop
	//		plot the pixel x, y with the color

	// For every pixel in the image you need two loops. One inside the other.
	// The outer loop goes top to bottom. The inner loop goes left to right
	//
	// You have to work top to bottom first so thats the y coordinate
	// You need to loop over the value of y while y is less than the window height.
	// remember y has to start at zero!
	// remember to make y bigger by one just before the end of the loop
	y = 0
	for y < windowHeight {
		// Now you have to work out the value of c that the pixel represents
		// We do this in two parts. First the imaginary part.
		// the current imaginary value of y is
		// maxIm - y * imFactor
		// see Owen for help with this bit!
		cIm = maxIm - float64(y)*imFactor

		// Now you have to work left to right. So you need to loop over the value of
		// x while x is less than the window width
		// remember x has to start at zero!
		// remember to make x bigger by one just before the end of the loop
		x = 0
		for x < windowWidth {
			// Now we need to work out the real part of c
			// the current real value of x is
			// minRe + x * reFactor
			// you need to use the same tip Owen gave you before
			cRe = minRe + float64(x)*reFactor

			// now we have to copy c to Z.
			// Easy - just two assignments.
			zRe = cRe
			zIm = cIm

			// Now we have to do the calculation (Z * Z) + c hundres of times!
			// You need another loop!
			// Set loopCount to zero
			loopCount = 0
			// loop while loopCount is less than maxLoopCount
			for loopCount < maxLoopCount {
				// The next bit is short cut to save you working these out again later
				// set zReSqrd to zRe * zRe
				zReSqrd = zRe * zRe
				// set zImSqrd to zIm * zIm
				zImSqrd = zIm * zIm

				// now do the escape test!
				// You can calculate the absolute value of Z by doing
				// zReSqrd + zImSqrd
				// If the answer is greater than 4 you have escaped!
				if zReSqrd+zImSqrd > 4 {
					// to stop the loop use the new break keyword. See Owen for help
					// we escaped! So this value of c is NOT in the set
					break
				}

				// You did not escape. So now you need to calculate the new value of Z
				// First you need to calculate the new value of Z's imaginary part
				// set zIm to 2 * zRe * zIm + cIm
				zIm = 2*zRe*zIm + cIm
				// now you need to calculate a new value of Z's real part
				// set zRe to zReSqrd - zImSqrd + cRe
				zRe = zReSqrd - zImSqrd + cRe
				// remember to make you loop counter one bigger here!
				loopCount = loopCount + 1
			}

			// We have escaped! The inner most loop has stopped becuase it reached
			// maxLoopCount or we escaped.
			// Now you have to work out the color based on the number of times you
			// went around the loop
			// You have to use the getColor function for this. You have to tell
			// getColor the number of times around the loop by putting loopCount
			// inside the brackets.You need to assign the result to
			// the variable called color. Like this
			// colour = getColor(loopCount)
			colour = getColor(loopCount)
			// Now you need to plot the point in the Window
			// You have to use the draw point function for this like this
			// drawPoint(x, y, color)
			drawPoint(x, y, colour)
			// now make x bigger here!
			x = x + 1
		}
		// now make y bigger here
		y = y + 1
		// uncomment the next line to see the image drawn one line at a time
		// Warning: The program will be a LOT slower!
		//render(x, y)
	}
	// This line draws the finished picture
	render(windowWidth, windowHeight)

	// Tell the user that you have finished
	fmt.Println("Finished")
	// now work out how long it took to draw the mandelbrot set.
	// Fist get the finish time
	var endTime time.Time
	endTime = time.Now()
	// now print out the differene between the start and end times.
	fmt.Printf("The program took %v to run.\n", endTime.Sub(startTime))
	// wait until you close the window before the program ends.
	waitUntilCloseButtonIsPressed()
}

//----------------- Don't change anythign below this line! ----------------

// Get a colour from the palette based on the number of iterations
func getColor(count int) toolbox.Colour {
	var c toolbox.Colour
	// try changing BluePalette to RedPalette to draw the mandelbrot in red
	c = palette.BluePalette[count%len(palette.BluePalette)]
	return c
}

// Plot the point x, y on window with the colour c
func drawPoint(x, y int, c toolbox.Colour) {
	plot[x][y].R = c.R
	plot[x][y].G = c.G
	plot[x][y].B = c.B
	plot[x][y].A = c.A
}

// Render or draw the picture to the window
func render(maxX, maxY int) {
	var x int
	var y int
	x = 0
	y = 0
	toolbox.SetBackgroundColour(0, 0, 0)
	toolbox.ClearBackground()
	for y < maxY {
		for x < maxX {
			toolbox.SetDrawColour(plot[x][y])
			toolbox.DrawPoint(x, y)
			x = x + 1
		}
		y = y + 1
		x = 0
	}
	toolbox.ShowWindow()
}

// Wait for the event that tells us that the user has pressed windows close
// button.
func waitUntilCloseButtonIsPressed() {
	var quit bool
	quit = false
	var key toolbox.Key

	for quit != true {
		key = toolbox.GetKey()
		if key == toolbox.ButtonClose {
			quit = true
		}
	}
}
