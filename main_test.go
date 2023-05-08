package main

import (
	"testing"
	"time"
)

func TestGetPokemons(t *testing.T) {
	start := time.Now()
	pokemons := GetPokemons()
	s := int(time.Since(start).Seconds())
	if s > 3 {
		t.Errorf("Ваш код работает слишком медленно (%v сек.); ожидается <= 3 сек.", s)
	}
	collected := len(pokemons)
	if collected < 1281 {
		t.Errorf("Собрано покемонов %v; ожидается 1281", collected)
	}
	for _, i := range pokemons {
		var x interface{} = i
		p, ok := x.(Pokemon)
		if !ok {
			t.Error("Неверный тип покемона", p)
		}
	}
}
