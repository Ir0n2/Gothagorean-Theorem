package main

import ("fmt"; "math")

func main () {

var A, B float64

var H float64

var shape string

fmt.Println("Type 'square' to calculate surface area of a square,\n 'Triangle' to calculate surface area of a triangle,\n 'break'to exit the program,\n and 'Hypotenuse' to calculate pythagorean theorem!\n or type help for more commands!")

looper:for {

        fmt.Scanln(&shape)

switch shape {

	case "square", "Square", "SQUARE":

fmt.Println("What is your squares length?")

        fmt.Scanf("%g", &A)

fmt.Println("What's the width of your square?")

        fmt.Scanf("%g", &B)

fmt.Println("the area of your square is ", A * B)

	case "triangle", "Triangle", "TRIANGLE":

fmt.Println("What is your triangles base length?")

        fmt.Scanf("%g", &A)

fmt.Println("What's the height of your triangle?")

        fmt.Scanf("%g", &B)

fmt.Println("the area of your triangle is ", A * B / 2)

	case "hypotenuse", "Hypotenuse", "HYPOTENUSE": 

fmt.Println("Type Leg A")

        fmt.Scanf("%g", &A)

fmt.Println("Type Leg B")

        fmt.Scanf("%g", &B)

var ans float64 = ((A * A) + (B * B))
result := math.Sqrt(ans)

fmt.Println(result)

//You can add more cases here!//
	case "circle", "Circle", "CIRCLE":

fmt.Println("What is the radius of your circle?")

	fmt.Scanf("%g", &A)

	sqrans := A * A

fmt.Println("The area of your circle is ", sqrans * 3.14)


	case "Cylinder", "cylinder", "CYLINDER":

fmt.Println("What is the radius of your cylinder?")

	fmt.Scanf("%g", &A)

fmt.Println("what is the height of your cylinder?")

	fmt.Scanf("%g", &B)

	sqrad := A * A

fmt.Println("The total surface area of your cylinder is", (2 * (3.14 * sqrad)) + ((A * B) * 6.28))

case "sphere", "SPHERE", "Sphere":

fmt.Println("What is the radius of your sphere?")

	fmt.Scanf("%g", &A)

fmt.Println("The surface area of your swuare is", 4 * (3.14 * (A * A)))

case "cube", "CUBE", "Cube":

fmt.Println("Enter The length of your cube/rectangle")

	fmt.Scanf("%g", &A)

fmt.Println("Enter The width of your cube/rectangle")

	fmt.Scanf("%g", &B)

fmt.Println("Enter The height of your cube/rectangle")

	fmt.Scanf("%g", &H)

fmt.Println("The area of your cube is:", (2 *(A * B)) + (2 * (A * H)) + (2 * (B * H)))

case "cone", "Cone", "CONE":

fmt.Println("Enter the radius of your cone")

	fmt.Scanf("%g", &A)

fmt.Println("Enter the slope of your cone")

	fmt.Scanf("%g", &B)

conesqr := ((A * A) * 3.14)

fmt.Println("The surface area of your cone is:", conesqr + (B * A * 3.14))


//reminder to document your code in this case//
	case "Help", "help", "HELP":

fmt.Println(" circle = area of circle\n cylinder = total surface area of a cylinder\n square = area of square\n triangle = area of triangle\n hypotenuse = finds the hypotenuse of a right triangle\n sphere = surface area of sphere\n cube = surface area of cube\n cone = surface area of a cube\n")

//reminder to document your code is this case//

case "break", "Break", "BREAK":

fmt.Println("breaking")

break looper

//just make sure to keep the default case at the bottem!//
default:
	fmt.Println("fuck off...")

		}
	}
}
