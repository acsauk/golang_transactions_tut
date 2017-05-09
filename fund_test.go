package funding

import (
  "sync"
  "testing"
  )

const WORKERS = 10

func BenchmarkWithdrawals(b *testing.B) {
  // Skip N = 1
  if b.N < WORKERS {
    return
  }

  // Add as many pounds as there are iterations in this run
  fund := NewFund(b.N)

  // Casually assume b.N divides cleanly
  poundsPerFounder := b.N / WORKERS

  // Waitgroup structs don't need to be initialized
  // (their zero value is ready to use)
  // Just declare one and then use it
  var wg sync.WaitGroup

  for i := 0; i < WORKERS; i++ {
    // Let the waitgroup know we're adding a goroutine
    wg.Add(1)

    //Spawn off a founder worker, as a closure
    go func() {
      //Mark this worker done when the function finishes
      defer wg.Done()

      for i := 0; i < poundsPerFounder; i++ {
        fund.Withdraw(1)
      }
    }() // The brackets make this a closure that has been called
  }

  //Wait for all the workers to finish
  wg.Wait()

  if fund.Balance() != 0 {
    b.Error("Balance wasn't zero:", fund.Balance())
  }

}

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

  func() {
    resource.Lock()
    defer resource.Unlock()

    // Do stuff with the resource here...
  }
}
