package person

func CreateArrayOfPeople() []Person {
	kinga := Person{Age: 20, Name: "Kinga"}
	filip := Person{Age: 20, Name: "Filip"}
	david := Person{Age: 20, Name: "Dawid"}
	karol := Person{Age: 19, Name: "Karol"}

	return []Person{kinga, david, filip, karol}
}
