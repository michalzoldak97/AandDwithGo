package decorator

type PerformanceTires struct {
	baseOffer Product
}

func (p *PerformanceTires) SetBaseOffer(pr Product) {
	p.baseOffer = pr
}

func (p *PerformanceTires) GetOffer() int {
	return p.baseOffer.GetOffer() + 7000
}
