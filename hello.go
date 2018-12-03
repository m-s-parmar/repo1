package main

import (
	"fmt"
	"errors"
	"math"
	"libx"
)

type person struct{
	name string
	age int
}

func main() {
	arr := []string {"a","b","c","d"}

	for index ,value := range arr{
		fmt.Println("index:",index,"value:",value)
	}


	result := sum(2,3)
	fmt.Println(result)


	sqrtResult,err := sqrt(16)
	
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(sqrtResult)
	}

	
	p1 := person{name:"Jake", age:23}
	fmt.Println(p1)

	p2 := person{"John",27}
	fmt.Println(p2.age)


	i := 7
	inc(&i)
	fmt.Println(i)


	x := true
	fmt.Println(x)


	language := libx.Get("go")
	
	fmt.Println(language)


	slice := []int{1,2,3,4,5}
	slice2 := make([]int, 5)
	copy(slice2,slice)

	for _,v := range slice2{
		fmt.Println(v)
	}


	defer printOne()
	

	num := 3
	doubleNum := func() int{
		return num * 2
	}

	fmt.Println(doubleNum())


	fmt.Println(safeDiv(3,0))
	fmt.Println(safeDiv(100,2))


	demPanic()


	yPtr := new(int)
	changeYVal(yPtr)
	fmt.Println("yPtr:",*yPtr)


	r := rectangle{10,20}
	fmt.Println("Area of rectangle:",r.recArea())


	sq := square{4}
	fmt.Println("Area of square:", getArea(&sq))

	c := circle{4}
	fmt.Println("Area of circle:", getArea(c))
}
//////////////////////////////////////////////////////////////


func sum(x int, y int) int{
	return x+y
}


func sqrt(x float64) (float64,error){
	 if x < 0{
		 return 0, errors.New("Undefined for a negative number")
	 }
	 return math.Sqrt(x),nil
}


func inc(x *int){
	*x++
}


func printOne(){
	fmt.Println(1)
}


func safeDiv(num1, num2 int) int{
	defer func(){
		fmt.Println(recover())
	}()

	solution := num1 / num2

	return solution
}


func demPanic(){
	defer func(){
		fmt.Println(recover())
	}()

	panic("PANIC")
}


func changeYVal(ptr *int){
	*ptr = 100
}


type rectangle struct{
	height float64
	width float64
}

func (rect *rectangle) recArea() float64{
	return rect.width * rect.height
}

//__________________________________

type shape interface{
	area() float64
}

type square struct{
	side float64
}

type circle struct{
	radius float64
}

func (sq *square) area() float64{
 	return sq.side * sq.side
}

func (circ circle) area() float64{
	return math.Pi * math.Pow(circ.radius,2)
}

func getArea(shape shape) float64{
	return shape.area()
}
//_________________________________