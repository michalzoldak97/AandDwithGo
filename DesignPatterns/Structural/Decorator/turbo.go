package decorator

type Turbo struct {
	baseOffer Product
}

func (t *Turbo) SetBaseOffer(pr Product) {
	t.baseOffer = pr
}

func (t *Turbo) GetOffer() int {
	return t.baseOffer.GetOffer() + 30000
}
