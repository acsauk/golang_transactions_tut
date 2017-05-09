package funding

import "testing"

func BenchmarkFund(b *testing.B) {
  // Add as many pounds as there are iterations in this run
  fund := NewFund(b.N)

  // Withdraw the deposit one pound at a time until the balance is 0
  for i := 0; i < b.N; i++ {
    fund.Withdraw(1)
  }

  if fund.Balance() != 0 {
    b.Error("Balance wasn't zero:", fund.Balance())
  }
}
