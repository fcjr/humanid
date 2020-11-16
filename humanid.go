package humanid

//go:generate go run gen.go

import (
	"math/rand"
	"strings"
)

var defaultProvider = &Provider{}

func New() string {
	return defaultProvider.New()
}

type Provider struct {
	sep string
}

func NewProvider(sep string) *Provider {
	return &Provider{
		sep: sep,
	}
}

func (p *Provider) New() string {
	noun := nouns[rand.Intn(len_nouns)]
	verb := verbs[rand.Intn(len_verbs)]
	adjective := adjectives[rand.Intn(len_adjectives)]
	sep := p.sep
	if p.sep == "" {
		sep = "-"
	}
	return strings.Join([]string{noun, verb, adjective}, sep)
}
