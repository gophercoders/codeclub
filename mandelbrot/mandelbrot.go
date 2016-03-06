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
	reFactor = (maxRe - minRe) / float64(windowWidth)
	imFactor = (maxIm - minIm) / float64(windowHeight)

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
	// Your Turn
	// first you have to set the y variable to zero
	// Write you answer on the line below

	// Your turn
	// You have to work top to bottom first. This is the outer loop
	// You have to loop while the value of y is less than the window  height
	// How do you do this?
	for  {
		// Now you have to work out the value of c that the pixel represents
		// We do this in two parts. First the imaginary part.

		// Your turn
		// You can work out current imaginary value, the cIm variable, of y like
		// this:
		// maxIm - y * imFactor
		// IMPORTANT
		// This answer isn't perfect you need to ask for help with this bit!
		// When you are left wth only two errors in the program ask for help
		// The first error will be here
		cIm =

		// Your Turn
		// First you have to set the value of x to zero
		// Write your answer on the line below

		// Now you need another loop to work left to right across the screen
		// Your Turn
		// You have to loop while the value of x is less than the window width
		// How do you do this?
		for  {
			// Now we need to work out the real part of c
			// Your turn
			// You can work out the current real value, the cRe value, of x
			// like this:
			// minRe + x * reFactor
			// IMPORTANT
			// This answer isn't perfect you need to ask for help with this bit!
			// When you are left wth only two errors in the program ask for help
			// The first error will be here
			cRe =

			// Now we have to copy c to Z.
			// This is easy - just two assignments.
			// Your turn
			// You need to set the value of zRe to cRe
			// then set the value of zIm = cIm

			// Now we have to do the calculation (Z * Z) + c hundres of times!
			// So we need a 3rd loop

			// Your turn
			// We use the variable called loopCount to count how many times the
			// program goes this loop.
			// First you have to set loopCount to zero.
			// Write your answer on the next line

			// Your Turn
			// Now you need to loop while loopCount is less than the maxLoopCount
			// How do you write this?
			for  {
				// The next bit is short cut to save you working these out again later
				// You have to work out the value of Z * Z
				// Your turn
				// You need to work out Z * Z in two steps
				// First you have to work out the real part, zReSqrd, by
				// multiplying zRe by zRe
				// How do you do this?
				zReSqrd =

				// Your Turn
				// Second you have to work out the imaginary part, zImSqrd, by
				// multiplying zIm by zIm
				// How do you do this?
				zImSqrd =

				// Now you need the escape test!
				//
				// Your Turn
				// The escape test is to test if the absolute value of Z is
				// greater than 4. If it is you have to escape by breaking the
				// 3rd loop.
				// You can calculate the absolute value of Z by doing
				// zReSqrd + zImSqrd
				// How do you write a condition that does this?
				if  {
					// to stop the loop use the new break keyword.
					// This value of c is NOT in the set
					break
				}

				// You did not escape. So now you need to calculate the new value of Z
				// Your Turn
				// First you need to calculate the new value of Z's imaginary part
				// zIm like this
				// 2 * zRe * zIm + cIm
				zIm =

				// Your Turn
				// Now you need to calculate a new value of Z's real part,
				// zRe like this
				// zReSqrd - zImSqrd + cRe
				zRe =

				// Your Turn
				// remember to make the loop counter one bigger here!
				loopCount = loopCount + 1
			}

			// We have escaped! The inner most (or 3rd) loop has stopped because it reached
			// maxLoopCount or we escaped.
			// Now you have to work out the color based on the number of times you
			// went around the loop

			// You have to use the getColor function for this. You have to tell
			// getColor the number of times around the loop by putting loopCount
			// inside the brackets.You need to assign the result to
			// the variable called color.
			colour = getColor(loopCount)
			// Now you need to plot the point in the Window
			// You have to use the draw point function for this
			drawPoint(x, y, colour)

			// Your turn
			// Now you have worked out the colour for this pixel you need to
			// work out the colour for the next pixel, moving horizontailly to
			// the right.
			// You do this by making x one bigger.
			// How do you do this?

		}
		// Your turn
		// Now you have worked out the colour for this pixel you need to
		// work out the colour for the next pixel, moving vertically downwards
		// You do this by making y one bigger.
		// How do you do this?


		// uncomment the next line to see the image drawn one line at a time
		// Warning: The program will be a LOT slower!
		//render(x, y)
	}
	// This line draws the finished picture
	// If you uncomment the render line above you need to commet out the
	// next line
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
