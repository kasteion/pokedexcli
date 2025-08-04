package pokeapi

import "sync"

type Pokedex struct {
	pokemons map[string]Pokemon
	mu *sync.Mutex
}

func newPokedex() Pokedex {
	p := Pokedex{
		pokemons: make(map[string]Pokemon),
		mu: &sync.Mutex{},
	}

	return p
}

func (p *Pokedex) Add(key string, pokemon Pokemon) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.pokemons[key] = pokemon
}