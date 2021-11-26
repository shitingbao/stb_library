package polish

import (
	"fmt"
	"strconv"
	"unicode"
)

func _example() {

	str := `5+7+6*(10+6)-1` // 基本数学表达式
	fmt.Println("表达式：" + str)

	arr := ToArr(str) //数组转化
	fmt.Println("toArr:", arr)

	re := Change(arr) // 波兰表达式
	fmt.Println("Change:", re)

	result := Js(re) // 表达式计算
	fmt.Println("Js:", result)

}

func ToArr(str string) []string {

	stackobject1 := new(ItemStack)
	stackResult := stackobject1.New()
	s := ""
	for i, v := range str {
		if v > 47 && v < 59 || v == 46 {
			s += string(v)
			if i == len(str)-1 {
				stackResult.Push(s)
			}
		} else {
			if s != "" {
				stackResult.Push(s)
				s = ""
			}
			stackResult.Push(string(v))
		}
	}

	return stackResult.Get()
}

func Change(str []string) []string {
	//48-57 46
	stackobject1 := new(ItemStack)
	stackResult := stackobject1.New()
	stackobject2 := new(ItemStack)
	stackYs := stackobject2.New()

	for _, j := range str {

		s := string(j)
		switch s {
		case "(":
			stackYs.Push(s)

		case ")":
			for !stackYs.IsEmpty() {
				preChar := stackYs.Top()
				if preChar == "(" {
					stackYs.Pop() // 弹出 "("
					break
				}
				stackResult.Push(preChar)
				stackYs.Pop()
			}

		case "+", "-", "*", "/":
			for !stackYs.IsEmpty() {

				if stackYs.Top() == "(" || isLower(stackYs.Top(), s) {
					break
				}
				stackResult.Push(stackYs.Pop())
			}
			stackYs.Push(s)

		default:
			stackResult.Push(s)
		}

	}

	for !stackYs.IsEmpty() {
		stackResult.Push(stackYs.Pop())
	}
	return stackResult.Get()
}

func Js(arr []string) []string {
	stackobject1 := new(ItemStack)
	stackResult := stackobject1.New()
	for _, v := range arr {
		switch v {
		case "+", "-", "*", "/":
			if stackResult.IsEmpty() {
				fmt.Println("error")
			}
			first, err := strconv.ParseFloat(stackResult.Pop(), 32)
			CheckErr(err)
			second, err := strconv.ParseFloat(stackResult.Pop(), 32)
			CheckErr(err)
			switch v {
			case "+":
				reStr := second + first
				s1 := strconv.FormatFloat(reStr, 'f', -1, 32)
				stackResult.Push(s1)
			case "-":
				reStr := second - first
				s1 := strconv.FormatFloat(reStr, 'f', -1, 32)
				stackResult.Push(s1)
			case "*":
				reStr := second * first
				s1 := strconv.FormatFloat(reStr, 'f', -1, 32)
				stackResult.Push(s1)
			case "/":
				reStr := second / first
				s1 := strconv.FormatFloat(reStr, 'f', -1, 32)
				stackResult.Push(s1)
			}
		default:
			stackResult.Push(v)
		}
	}
	return stackResult.Get()
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func isLower(top string, newTop string) bool {
	// 注意 a + b + c 的后缀表达式是 ab + c +，不是 abc + +
	switch top {
	case "+", "-":
		if newTop == "*" || newTop == "/" {
			return true
		}
	case "(":
		return true
	}
	return false
}

// func main() {
// 	log.Println(calculate([]string{}))
// }

func calculate(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			default:
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}

func calculate2(postfix string) int {
	stack := ItemStack{}
	fixLen := len(postfix)
	for i := 0; i < fixLen; i++ {
		nextChar := string(postfix[i])
		// 数字：直接压栈
		if unicode.IsDigit(rune(postfix[i])) {
			stack.Push(nextChar)
		} else {
			// 操作符：取出两个数字计算值，再将结果压栈
			num1, _ := strconv.Atoi(stack.Pop())
			num2, _ := strconv.Atoi(stack.Pop())
			switch nextChar {
			case "+":
				stack.Push(strconv.Itoa(num1 + num2))
			case "-":
				stack.Push(strconv.Itoa(num1 - num2))
			case "*":
				stack.Push(strconv.Itoa(num1 * num2))
			case "/":
				stack.Push(strconv.Itoa(num1 / num2))
			}
		}
	}
	result, _ := strconv.Atoi(stack.Top())
	return result
}
