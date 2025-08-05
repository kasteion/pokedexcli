package pokeapi

import (
	"sync"
)

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

func (p *Pokedex) Get(key string) (Pokemon, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	pokemon, ok := p.pokemons[key]
	return pokemon, ok
}

func (p *Pokedex) GetPokemonsChannel() (chan Pokemon, int) {
	ch := make(chan Pokemon)

	go func() {
		p.mu.Lock()
		for _, pokemon := range(p.pokemons) {
			ch <- pokemon
		}
		defer p.mu.Unlock()
	}()
	
	return ch, len(p.pokemons)
}