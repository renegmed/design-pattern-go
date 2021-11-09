package main

type otp struct {
	iOtp iOtp
	name string
}

func newOtp(iotp iOtp) *otp {
	return &otp{iOtp: iotp}
}

func (o *otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	o.iOtp.publishMetric()
	return nil
}

func (o *otp) setName(name string) {
	o.name = name
}

func (o *otp) getName() string {
	return o.name
}
