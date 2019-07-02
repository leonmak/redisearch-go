package redisearch

type Predicate struct {
	Property string
	// incluive
	Min string
	Max string
}

func (p *Predicate) Values() []string {
	max := p.Max
	min := p.Min
	if max == "" {
		max = "+inf"
	}
	if min == "" {
		min = "-inf"
	}
	return []string{p.Property, min, max}
}
