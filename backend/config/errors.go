package config

type Errors struct {
	errors []error
}

func (p *Errors) Add(err error) {
	p.errors = append(p.errors, err)
}

func (p *Errors) HasErrors() bool {
	return len(p.errors) > 0
}

func (p *Errors) ForEach(errFunc func(err error)) {
	for _, err := range p.errors {
		errFunc(err)
	}
}
