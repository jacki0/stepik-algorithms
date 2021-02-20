package main

/*Вы разрабатываете текстовый редактор для программистов и хотите реализовать проверку корректности расстановки скобок.
В коде могут встречаться скобки[]{}(). Из них скобки[,{и(считаются открывающими, а соответствующими им закрывающими
скобками являются],}и). В случае, если скобки расставлены неправильно, редактор должен также сообщить пользователю
первое место, где обнаружена ошибка.В первую очередь необходимо найти закрывающую скобку, для которой либо нет
соответствующей открывающей (например, скобка] в строке “]()”), либо же она закрывает не соответствующую ей открывающую
скобку (пример: “()[}”). Если таких ошибок нет, необходимо найти первую открывающую скобку, для которой нет
соответствующей закрывающей (пример: скобка(в строке “{}([]”). Помимо скобок, исходный код может содержать символы
латинского алфавита, цифры и знаки препинания.*/

import "fmt"

func Contains(a []string, x string) int {
	for i := len(a) - 1; i >= 0; i-- {
		if x == a[i] {
			return i
		}
	}
	return -1
}

func checkBrackets(sample string) int {
	var (
		brackets []string
		nums     []int
	)
	brack := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	for i := 1; i <= len(sample); i++ {
		brackets = append(brackets, sample[i-1:i])
		nums = append(nums, i)
	}
	for k := 0; len(brackets) > 1; {
		if k > 2 {
			break
		}
		for i, b := range brackets {
			if b == ")" || b == "]" || b == "}" {
				if i == 0 {
					return nums[0]
				}
				b = brack[b]
				n := Contains(brackets[:i], b)
				if n == -1 {
					return nums[i]
				} else {
					brackets = append(brackets[:i], brackets[i+1:]...)
					brackets = append(brackets[:n], brackets[n+1:]...)
					nums = append(nums[:i], nums[i+1:]...)
					nums = append(nums[:n], nums[n+1:]...)
					k = 0
					break
				}
			} else if b != "(" && b != "[" && b != "{" {
				brackets = append(brackets[:i], brackets[i+1:]...)
				nums = append(nums[:i], nums[i+1:]...)
				k = 0
				break
			}
		}
		k++
	}
	if len(nums) != 0 && len(brackets) != 0 {
		for i := len(nums) - 1; i >= 0; i-- {
			b := brackets[i]
			if b == ")" || b == "]" || b == "}" || b == "(" || b == "[" || b == "{" {
				return nums[i]
			}
		}
	}
	return -1
}

func main()  {
	input := []string{
		"[]([];",
		"({[",
		"({{}",
		"{",
		"{[]",
		"{{{",
		"[]([]",
		"{{{[][][]",
		"{{{{{{{((()))}",
		"{()}{",
		"}()",
		"()}()",
		"}()",
		"{[()]}}()",
		"dasdsadsadas]]]",
		"{}()",
		"({}[(((())))])",
		"()",
		"({})",
		"foo(bar({ <some initialization> })[i]);",
		"([](){([])})",
		"()[]}",
		"{{[()]]",
		"{{{[][][]",
		"{*{{}",
		"[[*",
		"{{",
		"{{{**[][][]",
	}
	tst := []string{
		"3",
		"3",
		"2",
		"1",
		"1",
		"3",
		"3",
		"3",
		"6",
		"5",
		"1",
		"3",
		"1",
		"7",
		"13",
		"Success",
		"Success",
		"Success",
		"Success",
		"Success",
		"Success",
		"5",
		"7",
		"3",
		"3",
		"2",
		"2",
		"3",
	}
	for i, sample := range input {
		res := checkBrackets(sample)
		if res == -1 {
			fmt.Print("Success")
		} else {
			fmt.Print(res)
		}
		fmt.Println(tst[i])
	}
}
