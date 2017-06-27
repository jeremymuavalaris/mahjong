package "mahjong"


type Hand struct {
	tiles []Tile
}

//GetNewHand returns an empty hand
func GetNewHand() *Hand{
	hand := Hand{}
	return *hand
}

//Adds a tile to the hand
func AddTile(hand *Hand, tile Tile) {

}

func OrderHand(hand *Hand) *Hand{
	if len(hand.tiles) < 2 { 
		return hand 
	}

	left, right := 0, len(hand.tiles) - 1

	// Pick a pivot
	pivotIndex := rand.Int() % len(hand.tiles)

	// Move the pivot to the right
	hand.tiles[pivotIndex], hand.tiles[right] = hand.tiles[right], hand.tiles[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range hand.tiles {
		if hand.tiles[i] < hand.tiles[right] {
			hand.tiles[i], hand.tiles[left] = hand.tiles[left], hand.tiles[i]
			left++
		}
		//If the tiles are of type circles, numbers, or sticks
		else if hand.tiles[i] == hand.tiles[right] {
			if //finish this later
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	qsort(a[:left])
	qsort(a[left + 1:])


	return a
}