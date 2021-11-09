package main

import "fmt"

func main() {
	// otp := otp{}

	// smsOTP := &sms{
	//  otp: otp,
	// }

	// smsOTP.genAndSendOTP(smsOTP, 4)

	// emailOTP := &email{
	//  otp: otp,
	// }
	// emailOTP.genAndSendOTP(emailOTP, 4)
	// fmt.Scanln()
	smsOTP := newSms("SMS-543 SYSTEM") //&sms{}
	// o := otp{
	// 	iOtp: smsOTP,
	// }
	o := newOtp(smsOTP)
	o.genAndSendOTP(4)

	fmt.Println("+++++++++++++++")

	emailOTP := newEmail("EMAIL-556 System") //&email{}
	// o = otp{
	// 	iOtp: emailOTP,
	// }
	o = newOtp(emailOTP)
	o.genAndSendOTP(4)

}
