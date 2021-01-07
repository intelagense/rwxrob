package lib

// -------------------------------------------------

type Person struct {
	Name string
	Age  int
}

func NewPerson() {
	person := Person{Name: "Rob", Age: 52}
	if person.Name == "Rob" {
	}
	if person.Age == 42 {
	}
}

// -------------------------------------------------

type PersonI interface {
	Name() string
	Age() int
}

type aperson struct {
	// fair to leave out since usually pointing to other data
}

func (p *aperson) Name() string {
	return "Rob"
}

func (p *aperson) Age() int {
	return 42
}

func NewPersonI() {
	person := new(aperson)
	if person.Name() == "Rob" {
	}
	if person.Age() == 42 {
	}
}

// -------------------------------------------------

type PersonII interface {
	Name() string
	Age() int
}

type aaperson struct {
	name string
	age  int
}

func (p *aaperson) Name() string {
	return p.name
}

func (p *aaperson) Age() int {
	return p.age
}

func NewPersonII() {
	person := aaperson{name: "Rob", age: 42}
	if person.Name() == "Rob" {
	}
	if person.Age() == 42 {
	}
}
