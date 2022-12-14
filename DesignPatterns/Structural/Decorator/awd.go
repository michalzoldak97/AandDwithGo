package decorator

type Awd struct {
	baseOffer Product
}

func (a *Awd) SetBaseOffer(pr Product) {
	a.baseOffer = pr
}

func (a *Awd) GetOffer() int {
	return a.baseOffer.GetOffer() + 15000
}
