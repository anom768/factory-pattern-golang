package employee

type Director struct {
	Employee
}

func NewDirector() *Director {
	// TODO: Create a new intern
	// Set the name to "Intern"
	// Set the salary to 100
	return &Director{
		Employee: Employee{
			Name:   "Director",
			Salary: 5000,
		},
	}
}

func (x *Director) GetBonus() float64 {
	return 0.3 * float64(x.Salary)
}
