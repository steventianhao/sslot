package engine

func MainGameReelStrips() [][]string {
	reels0 := [34]string{
		"Clam", "Nine", "Nemo", "Mermaid", "Queen", "Jack", "Green", "Ace", "Shark", "Starfish", "Ace",
		"King", "Clam", "Ten", "Queen", "Nemo", "Nine", "Starfish", "Mermaid", "Ten", "Starfish", "Queen",
		"Clam", "Octopus", "Ace", "Jack", "Mermaid", "King", "Queen", "Clam", "Ten", "Clam", "Ace", "Octopus",
	}
	reels1 := [33]string{
		"Green", "Ace", "Clam", "Mermaid", "King", "Octopus", "Ace", "Queen", "Starfish", "Shark", "Nemo",
		"King", "Starfish", "Mermaid", "Queen", "Nemo", "Jack", "Green", "Queen", "Ace", "Octopus", "King",
		"Green", "Ten", "Ace", "Starfish", "Nine", "King", "Green", "Ten", "Clam", "Nine", "Ten",
	}
	reels2 := [34]string{
		"King", "Jack", "Nemo", "Queen", "Octopus", "Ten", "Mermaid", "Ace", "Nemo", "Queen", "Starfish", "Ace",
		"King", "Starfish", "Ace", "Jack", "Octopus", "King", "Clam", "Shark", "Ten", "Nine", "Green", "Ace", "Nine",
		"Green", "Queen", "Mermaid", "Queen", "Octopus", "Ten", "Ace", "Clam", "Green",
	}
	reels3 := [46]string{
		"Jack", "Ten", "Starfish", "Shark", "Ten", "Green", "Nine", "Jack", "Octopus", "Mermaid", "Ace", "Jack",
		"Green", "Ace", "Nemo", "King", "Queen", "Clam", "King", "Jack", "Octopus", "Ace", "Nemo", "King", "Ace",
		"Starfish", "Nemo", "Queen", "King", "Nemo", "Nine", "Clam", "Ace", "King", "Nemo", "Ten", "Nine", "Octopus",
		"Jack", "King", "Starfish", "Ten", "Octopus", "Ace", "Nemo", "Ten",
	}
	reels4 := [35]string{
		"Nemo", "Starfish", "King", "Nine", "Mermaid", "Starfish", "Ace", "Nemo", "Octopus", "King", "Nine",
		"Starfish", "Mermaid", "Clam", "Queen", "Jack", "Octopus", "Ten", "Clam", "Ace", "Jack", "Green", "Octopus",
		"Nine", "Ten", "Shark", "King", "Octopus", "Queen", "Ace", "Clam", "Jack", "Octopus", "Queen", "Jack",
	}
	return [][]string{reels0[:], reels1[:], reels2[:], reels3[:], reels4[:]}
}

func AllSymbols() []*symbol {
	symbols := [13]*symbol{
		Ns(0, "Nine"),
		Ns(1, "Ten"),
		Ns(2, "Jack"),
		Ns(3, "Queen"),
		Ns(4, "King"),
		Ns(5, "Ace"),
		Ns(6, "Clam"),
		Ns(7, "Starfish"),
		Ns(8, "Nemo"),
		Ns(9, "Green"),
		Ns(10, "Octopus"),
		Ss(11, "Mermaid"),
		Ws(12, "Shark"),
	}
	return symbols[:]
}

func normalHits() map[HitKey]*Hit {
	hits := []*Hit{
		NewHit("Nine", 5, 100), NewHit("Nine", 4, 25), NewHit("Nine", 3, 5), NewHit("Nine", 2, 2),
		NewHit("Ten", 5, 100), NewHit("Ten", 4, 25), NewHit("Ten", 3, 5),
		NewHit("Jack", 5, 100), NewHit("Jack", 4, 25), NewHit("Jack", 3, 5),
		NewHit("Queen", 5, 125), NewHit("Queen", 4, 30), NewHit("Queen", 3, 5),
		NewHit("King", 5, 150), NewHit("King", 4, 40), NewHit("King", 3, 10),
		NewHit("Ace", 5, 150), NewHit("Ace", 4, 40), NewHit("Ace", 3, 10),
		NewHit("Clam", 5, 250), NewHit("Clam", 4, 75), NewHit("Clam", 3, 15),
		NewHit("Starfish", 5, 250), NewHit("Starfish", 4, 75), NewHit("Starfish", 3, 15),
		NewHit("Nemo", 5, 500), NewHit("Nemo", 4, 100), NewHit("Nemo", 3, 30),
		NewHit("Green", 5, 750), NewHit("Green", 4, 125), NewHit("Green", 3, 25), NewHit("Green", 2, 2),
		NewHit("Octopus", 5, 750), NewHit("Octopus", 4, 125), NewHit("Octopus", 3, 25), NewHit("Octopus", 2, 2),
	}
	return makeHitMap(hits)
}

func makeHitMap(hits []*Hit) map[HitKey]*Hit {
	m := make(map[HitKey]*Hit)
	for _, v := range hits {
		m[v.key()] = v
	}
	return m
}

func wildHits() map[HitKey]*Hit {
	hits := []*Hit{
		NewHit("Shark", 5, 10000),
		NewHit("Shark", 4, 2500),
		NewHit("Shark", 3, 250),
		NewHit("Shark", 2, 10),
	}
	return makeHitMap(hits)
}

func scatterHits() map[HitKey]*Hit {
	hits := []*Hit{
		NewFeatureHit("Mermaid", 5, 100, 15, 3),
		NewFeatureHit("Mermaid", 4, 10, 15, 3),
		NewFeatureHit("Mermaid", 3, 5, 15, 3),
	}
	return makeHitMap(hits)
}

func FeatureGameReelStrips() [][]string {
	reels0 := [35]string{
		"Clam", "Nine", "Nemo", "Mermaid", "Queen", "Jack", "Green", "Ace", "Shark", "Starfish", "Ace", "King",
		"Clam", "Shark", "Ten", "Queen", "Nemo", "Nine", "Starfish", "Mermaid", "Ten", "Starfish", "Queen", "Clam",
		"Octopus", "Ace", "Jack", "Mermaid", "King", "Queen", "Clam", "Ten", "Clam", "Ace", "Octopus",
	}
	reels1 := [34]string{
		"Green", "Ace", "Clam", "Mermaid", "King", "Octopus", "Ace", "Queen", "Starfish", "Shark", "Nemo", "King",
		"Starfish", "Mermaid", "Queen", "Nemo", "Jack", "Green", "Queen", "Shark", "Ace", "Octopus", "King", "Green",
		"Ten", "Ace", "Starfish", "Nine", "King", "Green", "Ten", "Clam", "Nine", "Ten",
	}
	reels2 := [35]string{
		"King", "Jack", "Nemo", "Queen", "Octopus", "Ten", "Mermaid", "Ace", "Nemo", "Queen", "Starfish", "Ace", "King",
		"Starfish", "Ace", "Shark", "Jack", "Octopus", "King", "Clam", "Shark", "Ten", "Nine", "Green", "Ace", "Nine",
		"Green", "Queen", "Mermaid", "Queen", "Octopus", "Ten", "Ace", "Clam", "Green",
	}
	reels3 := [48]string{
		"Jack", "Ten", "Starfish", "Shark", "Ten", "Green", "Nine", "Jack", "Octopus", "Mermaid", "Ace", "Jack", "Green",
		"Ace", "Nemo", "King", "Queen", "Clam", "King", "Shark", "Jack", "Octopus", "Ace", "Nemo", "King", "Shark", "Ace",
		"Starfish", "Nemo", "Queen", "King", "Nemo", "Nine", "Clam", "Ace", "King", "Nemo", "Ten", "Nine", "Octopus", "Jack",
		"King", "Starfish", "Ten", "Octopus", "Ace", "Nemo", "Ten",
	}
	reels4 := [37]string{
		"Nemo", "Starfish", "King", "Nine", "Mermaid", "Starfish", "Ace", "Nemo", "Octopus", "King", "Nine", "Starfish",
		"Mermaid", "Clam", "Queen", "Jack", "Shark", "Octopus", "Ten", "Clam", "Ace", "Jack", "Shark", "Green", "Octopus",
		"Nine", "Ten", "Shark", "King", "Octopus", "Queen", "Ace", "Clam", "Jack", "Octopus", "Queen", "Jack",
	}
	return [][]string{reels0[:], reels1[:], reels2[:], reels3[:], reels4[:]}
}

func lines() [][]int {
	return lines3x5
}
