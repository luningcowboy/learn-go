package main


import "fmt"

func main() {
	var grade string = "B"
	var marks int = 50

	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50, 60, 70: grade = "C"
		default: grade = "D"
	}

	switch {
		case grade == "A":
			fmt.Println("优秀\n")
		case grade == "B":
			fmt.Println("良好\n")
		case grade == "C":
			fmt.Println("及格\n")
		case grade == "D":
			fmt.Println("差\n")
		default:
			fmt.Println("啦啦啦\n")
	}
	fmt.Printf("你的等级是 %s \n", grade)

	// Type Switch
	var x interface{}
	switch i := x.(type){
	case nil:
		fmt.Printf("x 的类型: %T", i)
	case int:
		fmt.Printf("x 是int类型")
	case float64:
		fmt.Printf("x 是float64类型")
	case func(int) float64:
		fmt.Printf("x 是func(int)类型")
	case bool, string:
		fmt.Printf("x 是bool/string类型")
	default:
		fmt.Printf("未知类型")
	}

	println()
	// fallthrough
	switch{
	case false:
		fmt.Println("1. case = false")
		fallthrough
	case true:
		fmt.Println("2. case = true")
		fallthrough
	case false:
		fmt.Println("3. case = false")
		fallthrough
	case true:
		fmt.Println("4. case = true")
	case false:
		fmt.Println("5. case = false")
		fallthrough
	default:
		fmt.Println("6. 默认case")
	}
}
