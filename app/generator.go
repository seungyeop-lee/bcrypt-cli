package app

import "golang.org/x/crypto/bcrypt"

type Generator struct {
	cost int
}

func (g Generator) Cost() int {
	if g.cost == 0 || g.cost < bcrypt.MinCost || g.cost > bcrypt.MaxCost {
		return bcrypt.DefaultCost
	} else {
		return g.cost
	}
}

func NewGenerator(cost int) *Generator {
	return &Generator{cost: cost}
}

func (g Generator) Generate(password string) (string, error) {
	generated, err := bcrypt.GenerateFromPassword([]byte(password), g.Cost())
	return string(generated), err
}
