package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Avarage() float64 {
	sum := 0

	for _, score := range s.score {
		sum += score
	}

	return float64(sum) / float64(len(s.score))
}

func (s Student) Min() (min int, name string) {
	if len(s.score) == 0 {
		return 0, ""
	}

	min = s.score[0]
	name = s.name[0]

	for i, score := range s.score {
		if score < min {
			min = score
			name = s.name[i]
		}
	}

	return min, name
}

func (s Student) Max() (max int, name string) {
	if len(s.score) == 0 {
		return 0, ""
	}

	max = s.score[0]
	name = s.name[0]

	for i, score := range s.score {
		if score > max {
			max = score
			name = s.name[i]
		}
	}

	return max, name
}

func main() {
	var a = Student{}

	for i := 1; i < 6; i++ {
		var name string
		fmt.Print("Input ", i, " Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAvarage Score Students is", a.Avarage())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+"(", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+"(", scoreMin, ")")
}
