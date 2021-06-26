package lessons

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	numMatrix := NumMatrix{}
	for i := 0; i < len(matrix); i++ {
		numMatrix.sum = append(numMatrix.sum, []int{})
		for j := 0; j < len(matrix[i]); j++ {
			numMatrix.sum[i] = append(numMatrix.sum[i],
				numMatrix.getSum(i-1, j)+numMatrix.getSum(i, j-1)-numMatrix.getSum(i-1, j-1)+matrix[i][j])
		}
	}

	return numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.getSum(row2, col2) - this.getSum(row1-1, col2) - this.getSum(row2, col1-1) + this.getSum(row1-1, col1-1)
}

func (this *NumMatrix) getSum(i, j int) int {
	if i >= 0 && j >= 0 {
		return this.sum[i][j]
	}

	return 0
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
