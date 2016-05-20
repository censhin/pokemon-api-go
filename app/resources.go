package main

import (
    net/http
    log
)

func PokemonCollectionResource(w http.ResponseWriter, r *http.Request) {
    b, err := GetPokemonService()

    if err != nil {
        log.Fatal(err)
    } else {
        w.Write(b)
    }
}
