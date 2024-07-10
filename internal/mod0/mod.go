package mod0

import "blog/internal/cerror"

type M struct {
	name string
	age  int
}

func NewM(name string, age int) *M {
	return &M{
		name: name,
		age:  age,
	}
}
func (m *M) Info() *M {
	return m
}

func (m *M) GetName() string {
	return m.name
}

func (m *M) GetAge() int {
	return m.age
}

func (m *M) SetName(name string) error {
	m.name = name
	return nil
}

func (m *M) SetAge(age int) error {
	if age < 1 {
		return cerror.NewError("年龄不得小于1")
	}
	m.age = age
	return nil
}
