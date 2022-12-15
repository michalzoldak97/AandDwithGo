package decorator

type AWD struct {
	baseOffer Product
}

func (a *AWD) SetBaseOffer(pr Product) {
	a.baseOffer = pr
}

func (a *AWD) GetOffer() int {
	return a.baseOffer.GetOffer() + 15000
}
