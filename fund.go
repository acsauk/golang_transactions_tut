package funding

type Fund struct {
  //Lowercase 'balance' initializes a struct with an unexported (private) balance.
  balance int
}

func NewFund(initialBalance int) *Fund {
  // Its possible to return a 'pointer' to a new struct without the need
  // for it to be on the stack as Go works this out for us
  return &Fund{
    balance: initialBalance,
  }
}

// Methods being with a 'receiver' - for this case a Fund pointer
func (f *Fund) Balance() int {
  return f.balance
}

func (f *Fund) Withdraw(amount int) {
  f.balance -= amount
}
