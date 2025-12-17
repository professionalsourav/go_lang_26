package main

import "fmt"

func main() {

	//section1: intigers
	//section-A: signed integers
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	//assigning values to integer variable
	i =-128
	i8 = 127
	i16 = -32768
	i32 = 444464665
	i64 = -86756579767676888

	
	//section-B: unsigned integers
	var u uint
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64

	//assigning values to integer variable
	u = 188
	u8 = 127
	u16 = 32768
	u32 = 444464665
	u64 = 86756579767676888

	//printing the signed and unsigned integer variables to the console

	fmt.Println("signed intiger:", i, i8, i16, i32, i64);
	fmt.Println("unsigned intiger:", u, u8, u16, u32, u64);


	//section2: floating point
	//section2-A : float32
	//this is use for less precise calculation
     var f32 float32 = 10.6
	//section2-B : float32
     var f64 float64 = 10.6
	 //this is use for more precise calculation

	 //printing the flaoting point value
	 fmt.Println("float32:", f32)
	 fmt.Println("float64:", f64)

	 //Demonstrating the diff in precesion between the f32 and f64

	 var HP float64 = 10123456789012345

	  var LP float32 = 10123456789012345

	  fmt.Println("High precision float64:", HP)
	  fmt.Println("High precision float64:", LP)

	  //section3: Boolean datatype

	  var isActive bool = true;
	var isOn bool = false;

	fmt.Println("Is active:", isActive);
	fmt.Println("Is on:", isOn);

	//section4: Complex Type

	var CN1 complex128 = complex(5,10)
	var CN2 complex64 = complex(2,7)
	// print(CN1)
	// print(CN2)
	fmt.Println("CN1:", CN1);
	fmt.Println("CN2", CN2);

     //section5: string data type
	 var name string = "sourav is a good boy";
	 fmt.Println("string value:", name);
}
























