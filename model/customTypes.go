package model

import(
	"time"
)

type Node struct{
	Chrctr Character 
	Next Node
	Prev Node
}


type MarvelData struct{
	Last_update time.Time
	Size int
	CharacterMap map[string]Node
}

type Character struct{
	Name string
	Max_power int
} 

