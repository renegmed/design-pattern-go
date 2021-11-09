package steps

import (
	"chain_of_responsiblity_design_pattern/customer"
)

type step interface {
	Run(*customer.Customer)
	SetNextStep(step)
}
