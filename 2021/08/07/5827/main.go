package main

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	tempR := rMove
	if tempR < 6 && board[tempR+1][cMove] != color && board[tempR+1][cMove] != '.' {
		tempR += 2
		for tempR <= 7 {
			if board[tempR][cMove] == '.' {
				break
			}
			if board[tempR][cMove] == color {
				return true
			}
			tempR++
		}
	}

	tempR = rMove
	if tempR > 1 && board[tempR-1][cMove] != color && board[tempR-1][cMove] != '.' {
		tempR -= 2
		for tempR >= 0 {
			if board[tempR][cMove] == '.' {
				break
			}
			if board[tempR][cMove] == color {
				return true
			}
			tempR--
		}
	}
	tempC := cMove
	if tempC < 6 && board[rMove][tempC+1] != color && board[rMove][tempC+1] != '.' {
		tempC += 2
		for tempC <= 7 {
			if board[rMove][tempC] == '.' {
				break
			}
			if board[rMove][tempC] == color {
				return true
			}
			tempC++
		}
	}
	tempC = cMove
	if tempC > 1 && board[rMove][tempC-1] != color && board[rMove][tempC-1] != '.' {
		tempC -= 2
		for tempC >= 0 {
			if board[rMove][tempC] == '.' {
				break
			}
			if board[rMove][tempC] == color {
				return true
			}
			tempC--
		}
	}
	tempR = rMove
	tempC = cMove
	if tempC > 1 && tempR > 1 && board[tempR-1][tempC-1] != color && board[tempR-1][tempC-1] != '.' {
		tempC -= 2
		tempR -= 2
		for tempC >= 0 && tempR >= 0 {
			if board[tempR][tempC] == '.' {
				break
			}
			if board[tempR][tempC] == color {
				return true
			}
			tempC--
			tempR--
		}
	}
	tempR = rMove
	tempC = cMove
	if tempR < 6 && tempC < 6 && board[tempR+1][tempC+1] != color && board[tempR+1][tempC+1] != '.' {
		tempR += 2
		tempC += 2
		for tempR <= 7 && tempC <= 7 {
			if board[tempR][tempC] == '.' {
				break
			}
			if board[tempR][tempC] == color {
				return true
			}
			tempR++
			tempC++
		}
	}
	tempR = rMove
	tempC = cMove
	if tempR < 6 && tempC > 1 && board[tempR+1][tempC-1] != color && board[tempR+1][tempC-1] != '.' {
		tempR += 2
		tempC -= 2
		for tempR <= 7 && tempC >= 0 {
			if board[tempR][tempC] == '.' {
				break
			}
			if board[tempR][tempC] == color {
				return true
			}
			tempR++
			tempC--
		}
	}
	tempR = rMove
	tempC = cMove
	if tempR > 1 && tempC < 6 && board[tempR-1][tempC+1] != color && board[tempR-1][tempC+1] != '.' {
		tempR -= 2
		tempC += 2
		for tempR >= 0 && tempC <= 7 {
			if board[tempR][tempC] == '.' {
				break
			}
			if board[tempR][tempC] == color {
				return true
			}
			tempR--
			tempC++
		}
	}
	return false
}
