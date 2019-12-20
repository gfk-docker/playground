package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

// MockCompany is an autogenerated mock type for the Company type
type MockCompanyManually struct {
	mock.Mock
}

// Info provides a mock function with given fields: name
func (_m *MockCompanyManually) Info(name string) (string, string, string, int) {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 string
	if rf, ok := ret.Get(2).(func(string) string); ok {
		r2 = rf(name)
	} else {
		r2 = ret.Get(2).(string)
	}

	var r3 int
	if rf, ok := ret.Get(3).(func(string) int); ok {
		r3 = rf(name)
	} else {
		r3 = ret.Get(3).(int)
	}

	return r0, r1, r2, r3
}

func TestFoo_DoSomething(t *testing.T) {
	mockCompany := new(MockCompanyManually)
	mockCompany.On("Info", "Aqua").Return(
		"100-0000", "Ramat Gan", "+972-3-688-8799", 200)

	f := Foo{
		Company: mockCompany,
	}
	f.DoSomething()
}

func TestFoo_DoSomething_WithTableDrivenTest(t *testing.T) {
	type InfoArgs struct {
		CompanyName string
	}
	type InfoReturns struct {
		Zip      string
		Address  string
		Phone    string
		Employee int
	}
	type InfoExpectation struct {
		Args    InfoArgs
		Returns InfoReturns
	}

	tests := []struct {
		name string
		info InfoExpectation
	}{
		{
			name: "happy path",
			info: InfoExpectation{
				Args: InfoArgs{
					CompanyName: "Aqua",
				},
				Returns: InfoReturns{
					Zip:      "101-0000",
					Address:  "Ramat Gan",
					Phone:    "+972-3-688-8799",
					Employee: 2000,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCompany := new(MockCompanyManually)
			mockCompany.On("Info", tt.info.Args.CompanyName).Return(tt.info.Returns.Zip,
				tt.info.Returns.Address, tt.info.Returns.Phone, tt.info.Returns.Employee)

			f := Foo{
				Company: mockCompany,
			}
			f.DoSomething()
		})
	}
}

func TestFoo_DoSomething_WithMockery(t *testing.T) {
	tests := []struct {
		name string
		info InfoExpectation
	}{
		{
			name: "happy path",
			info: InfoExpectation{
				Args: InfoArgs{
					Name: "Aqua",
				},
				Returns: InfoReturns{
					Zip:      "101-0000",
					Address:  "Ramat Gan",
					Phone:    "+972-3-688-8799",
					Employee: 2000,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCompany := new(MockCompany)
			mockCompany.ApplyInfoExpectation(tt.info)
			f := Foo{
				Company: mockCompany,
			}
			f.DoSomething()
		})
	}
}