package lessons

func isValidSudoku(board [][]byte) bool {
	row, col, box := make([]map[byte]bool, 9), make([]map[byte]bool, 9), make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		row[i], col[i], box[i] = make(map[byte]bool, 0), make(map[byte]bool, 0), make(map[byte]bool, 0)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				digit := board[i][j]
				boxId := i/3*3 + j/3

				if _, ok := row[i][digit]; ok {
					return false
				}
				row[i][digit] = true

				if _, ok := col[j][digit]; ok {
					return false
				}
				col[j][digit] = true

				if _, ok := box[boxId][digit]; ok {
					return false
				}
				box[boxId][digit] = true

			}
		}
	}

	return true
}
