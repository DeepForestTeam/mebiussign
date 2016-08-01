package serviceproviders

import "github.com/DeepForestTeam/mobiussign/components/store"

func (this *ServiceProviderRow)AddService() (err error) {
	return
}

func (this *ServiceProviderRow)Validate() (err error) {
	return
}

func (this *ServiceProviderRow)validateLogin() (err error) {
	return
}

func (this *ServiceProviderRow)Register() (err error) {
	_, err = store.Set(ServicesStorage, "test", this)
	return
}