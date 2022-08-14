package main

import "fmt"

type ChainedExecutor interface {
	ExecuteNext(map[string]string)
}

type FullNameLogger struct {
	Next ChainedExecutor
}

func (f *FullNameLogger) ExecuteNext(data map[string]string) {
	firstName := data["firstName"]
	lastName := data["lastName"]
	fmt.Printf("My fullname is %s %s\n", firstName, lastName)

	if f.Next != nil {
		f.Next.ExecuteNext(data)
	}
}

type AddressLogger struct {
	Next ChainedExecutor
}

func (f *AddressLogger) ExecuteNext(data map[string]string) {
	city := data["city"]
	street := data["street"]
	country := data["country"]
	fmt.Printf("My address is %s in %s, %s\n", street, city, country)

	if f.Next != nil {
		f.Next.ExecuteNext(data)
	}
}

type BirthdayLogger struct {
	Next ChainedExecutor
}

func (f *BirthdayLogger) ExecuteNext(data map[string]string) {
	date := data["birthdate"]
	month := data["birthmonth"]
	fmt.Printf("My birthday is on %s-%s\n", date, month)

	if f.Next != nil {
		f.Next.ExecuteNext(data)
	}
}

func main() {
	data := map[string]string{
		"firstName":  "John",
		"lastName":   "Blake",
		"city":       "Doha",
		"country":    "Qatar",
		"street":     "Matr Qadeem Street No.50",
		"birthdate":  "19",
		"birthmonth": "March",
	}

	birthdayLogger := &BirthdayLogger{}
	addressLogger := &AddressLogger{Next: birthdayLogger}
	startingChainExecutor := &FullNameLogger{Next: addressLogger}
	startingChainExecutor.ExecuteNext(data)
}
