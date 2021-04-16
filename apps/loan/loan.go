package loan

import "time"

type Loan struct {
	IdUser       int
	IdBook       int
	BorrowedBook time.Time
	DueDate      time.Time
	ReturnDate   time.Time
}

type FinePayment struct {
	Receipt string
	Amount  float64
	IdLoan  int
}
