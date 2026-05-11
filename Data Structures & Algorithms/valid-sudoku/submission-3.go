func isValidSudoku(board [][]byte) bool {
	row:=map[int]map[byte]bool{}
	col:=map[int]map[byte]bool{}
	cel:=map[int]map[byte]bool{}
	for i:=0;i<9;i++{
		row[i] = map[byte]bool{}
		col[i] = map[byte]bool{}
		cel[i] = map[byte]bool{}
}
		for i:=0;i<9;i++{
			for j:=0;j<9;j++{
				val:=board[i][j]

				box:=(i/3)*3 + j/3
				if val=='.'{
					continue
				}
				if row[i][val]||col[j][val]||cel[box][val]{
					return false
				}
				row[i][val] = true
				col[j][val] = true
				cel[box][val] = true

			}
		}
		return true
}
