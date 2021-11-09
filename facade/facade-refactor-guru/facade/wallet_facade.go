package facade

import (
	"facade_design_pattern/facade/account"
	"facade_design_pattern/facade/ledger"
	"facade_design_pattern/facade/notification"
	"facade_design_pattern/facade/security"
	"facade_design_pattern/facade/wallet"
	"fmt"
)

type walletFacade struct {
	account      *account.Account
	wallet       *wallet.Wallet
	securityCode *security.SecurityCode
	notification *notification.Notification
	ledger       *ledger.Ledger
}

func NewWalletFacade(accountID string, code int) *walletFacade {
	fmt.Println("Starting create account")
	walletFacacde := &walletFacade{
		account:      account.NewAccount(accountID),
		securityCode: security.NewSecurityCode(code),
		wallet:       wallet.NewWallet(),
		notification: notification.NewNotification(),
		ledger:       ledger.NewLedger(),
	}
	fmt.Println("Account created")
	return walletFacacde
}

func (w *walletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.CreditBalance(amount)
	w.notification.SendWalletCreditNotification()
	w.ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

func (w *walletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.SendWalletDebitNotification()
	w.ledger.MakeEntry(accountID, "credit", amount)
	return nil
}
