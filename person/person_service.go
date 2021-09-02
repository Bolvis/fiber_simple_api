package person

func CreateArrayOfPeople() []Person {
	kinga := Person{Age: 20, Salary: 0, Name: "Kinga"}
	filip := Person{Age: 20, Salary: 3800, Name: "Filip"}
	david := Person{Age: 20, Salary: 6500, Name: "Dawid"}
	karol := Person{Age: 19, Salary: 7500, Name: "Karol"}

	return []Person{kinga, david, filip, karol}
}
