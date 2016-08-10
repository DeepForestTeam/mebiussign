package sign

import "github.com/DeepForestTeam/mobiussign/components/store"

func (this *MobiusSigner)Check(signature string) (err error) {
	if len(signature) == 16 {
		short_index := SignatureShortIndex{}
		err = store.Get(MobiusShortIndexStorage, signature, &short_index)
		if err == store.ErrKeyNotFound || err == store.ErrSectionNotFound {
			err = ErrSignNotFound
			return
		}
		signature = short_index.MobiusSign
	}
	err = store.Get(MobiusStorage, signature, &this.SignRow)
	if err == store.ErrKeyNotFound || err == store.ErrSectionNotFound {
		err = ErrSignNotFound
		return
	}
	this.fillResponse()
	return nil
}