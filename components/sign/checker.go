package sign

import "github.com/DeepForestTeam/mobiussign/components/store"

func (this *MobiusSigner)Check(signature string) (err error) {
	err = store.Get(MobiusStorage, signature, &this.SignRow)
	if err == store.ErrKeyNotFound || err == store.ErrSectionNotFound {
		err = ErrSignNotFound
		return
	}
	this.fillResponse()
	return nil
}