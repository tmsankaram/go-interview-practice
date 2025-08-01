package main

import "fmt"

type Employee struct {
  ID     int
  Name   string
  Age    int
  Salary float64
}

type Manager struct {
  Employees []Employee
}

// AddEmployee adds a new employee to the manager's list.
func (m *Manager) AddEmployee(e Employee) {
  // TODO: Implement this method
  m.Employees = append(m.Employees, e)
}

// RemoveEmployee removes an employee by ID from the manager's list.
func (m *Manager) RemoveEmployee(id int) {
  // TODO: Implement this method
  for i, emp := range m.Employees {
    if emp.ID == id {
      m.Employees = append(m.Employees[:i], m.Employees[i+1:]...)
      return
    }
  }
}

// GetAverageSalary calculates the average salary of all employees.
func (m *Manager) GetAverageSalary() float64 {
  // TODO: Implement this method
  // Check if the employees slice is empty
  if m == nil || len(m.Employees) == 0 {
    return 0
  }

  if len(m.Employees) == 0 {
    return 0
  }

  // Calculate the total sum of all salaries
  total := 0.0
  for _, employee := range m.Employees {
    total += employee.Salary
  }

  // Calculate and return the average
  return total / float64(len(m.Employees))
}

// FindEmployeeByID finds and returns an employee by their ID.
func (m *Manager) FindEmployeeByID(id int) *Employee {
  // TODO: Implement this method
  if m == nil {
    return nil
  }

  // Iterate through all employees
  for i := range m.Employees {
    if m.Employees[i].ID == id {
      // Return a pointer to the found employee
      return &m.Employees[i]
    }
  }

  // Return nil if employee not found
  return nil
}

func main() {
  manager := Manager{}
  manager.AddEmployee(Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
  manager.AddEmployee(Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})
  manager.RemoveEmployee(1)
  averageSalary := manager.GetAverageSalary()
  employee := manager.FindEmployeeByID(2)

  fmt.Printf("Average Salary: %f\n", averageSalary)
  if employee != nil {
    fmt.Printf("Employee found: %+v\n", *employee)
  }
}
