package pointers_and_errors

import (
	. "github.com/quii/pepper"
	. "github.com/quii/pepper/matchers/comparable"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := &Wallet{}
		wallet.Deposit(10)
		Expect(t, wallet).To(HaveBalance(EqualTo(10)))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := &Wallet{balance: 10}
		ExpectNoError(t, wallet.Withdraw(5))
		Expect(t, wallet).To(HaveBalance(EqualTo(5)))
	})

	t.Run("withdrawing too much returns an error", func(t *testing.T) {
		wallet := &Wallet{balance: 10}
		err := wallet.Withdraw(100)
		ExpectError(t, err)
		ExpectErrorOfType(t, err, ErrInsufficientFunds)
		Expect(t, wallet).To(HaveBalance(EqualTo(10)))
	})
}

func HaveBalance(matcher Matcher[int]) Matcher[*Wallet] {
	return func(wallet *Wallet) MatchResult {
		res := matcher(wallet.balance)
		res.SubjectName = "the balance of the wallet"
		return res
	}
}
