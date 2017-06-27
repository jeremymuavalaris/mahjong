package mahjong

//Tile is determined by type and number. Move to own file, when Game is created
//Ex. Six of sticks would be tile_type: 's', number: '6'
//'l' = sticks, 'c' = circles, 'd' = numbers
//'n' = north, 'w' = west, 'e' = east, 's' = south
//'g' = green dragon, 'r' = red dragon, 'b' = white dragon
type Tile struct {
	Value uint8
	TileType byte
}

var validTiles := map[byte]string {
	'l' : "sticks",
	'c' : "circles",
	'd' : "numbers",
	'n' : "north_direction",
	'w' : "west_direction",
	'e' : "east_direction",
	's' : "south_direction",
	'g' : "green_dragon",
	'r' : "red_dragon",
	'b' : "white_dragon"
}

//Creates a new tile after checking that it is valid, otherwise returns nil
func CreateTile(val uint8, tileType byte) *Tile {
	tile := &Tile{val, tileType}
	if IsValidTile(*tile) {
		return tile
	}
	*tile = nil
	return nil
}

//Check to see if tile is a valid circle, number, or stick
func IsNonHonorTile(tile Tile) bool {
	if (tile.Value != nil) && 
	(tile.TileType == 'l' || tile.TileType == 'c' || tile.TileType == 'd') &&
	(1 <= tile.Value && tile.Value <= 9) {
		return true;
	}
	return false;
}

//Checks to see if tile is a valid honor tile (wind or dragon)
func IsHonorTile(tile Tile) bool {
	if(tile.TileType == 'n' || tile.TileType == 'w' || tile.TileType == 'e' || tile.TileType == 's' ||
		tile.TileType == 'g' || tile.TileType == 'r' || tile.TileType == 'b') {
		return true
	}
	return false;
}


//Validates if the tile is valid
func IsValidTile(tile Tile) bool{
	if validTiles[tile.TileType] != nil {
		//check if tile is circle, number, or stick and if it has a valid value (1-9)
		if IsNonHonorTile(tile) || IsHonorTile(tile) {
			return true
		}
	}
	return false
}