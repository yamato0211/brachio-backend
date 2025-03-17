package model

type Monster struct {
	ID        string
	HP        int
	Elements  []Element
	SpawnedAt int // 召喚されたターン
}

type Element string
