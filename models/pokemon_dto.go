package models

type PokemonDTO struct {
	Name      string `json:"name" binding:"required,min=2,max=30,alphanum"`
	EvolvesAt int    `json:"evolves_at" binding:"required,min=1,max=100"`
}

type PokemonUpdateDTO struct {
	Name      string `json:"name,omitempty" binding:"omitempty,min=2,max=30,alphanum"`
	EvolvesAt int    `json:"evolves_at,omitempty" binding:"omitempty,min=1,max=100"`
}

func (p *Pokemon) ToDTO() PokemonDTO {
	return PokemonDTO{
		Name:      p.Name,
		EvolvesAt: p.EvolvesAt,
	}
}

func ToDTOs(pokemons []Pokemon) []PokemonDTO {
	dtos := make([]PokemonDTO, len(pokemons))
	for i, p := range pokemons {
		dtos[i] = p.ToDTO()
	}
	return dtos
}
