package flip_image

func flipAndInvertImage(A [][]int) [][]int {
	lenY := len(A)
	lenX := len(A[0])
	halfX := lenX / 2
	if lenX%2 == 1 {
		halfX += 1
	}
	for idxY := 0; idxY < lenY; idxY++ {
		for idxX := 0; idxX < halfX; idxX++ {
			A[idxY][idxX], A[idxY][lenX-1-idxX] = (A[idxY][lenX-1-idxX]+1)%2, (A[idxY][idxX]+1)%2
		}
	}
	return A
}
