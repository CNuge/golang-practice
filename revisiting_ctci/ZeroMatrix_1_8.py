"""
take an m x n matrix and if there is a zero in the matrix then change the entire 
row and column it is in to 0s
"""
import unittest

def ZeroMatrix(mat):
	""" take a list of lists where cells are integers,
		if there is a zero cell then change all values in the row and
		column to zero """

	z_rows = []
	z_cols = []

	for r, row in enumerate(mat):
		for c, val in enumerate(row):
			if val == 0:
				z_rows.append(r)
				z_cols.append(c)

	for r in set(z_rows):
		for i in range(0, len(mat[0])):
			mat[r][i] = 0

	for c in set(z_cols):
		for i in range(0,len(mat)):
			mat[i][c] = 0

	return mat


class TestZeroMat(unittest.TestCase):
	def test_ZeroMatrix(self):
		self.t1 = [[0,1,1,],
				[1,1,1,],
				[1,0,1,],
				[1,1,1,],]

		self.o1 = [[0,0,0,],
				[0,0,1,],
				[0,0,0,],
				[0,0,1,],]

		self.t2 = [[1,1,1,1,],
				[1,1,1,1,],
				[1,1,1,0,],
				[1,1,1,1,],]

		self.o2 = [[1,1,1,0,],
				[1,1,1,0,],
				[0,0,0,0,],
				[1,1,1,0,],]

		self.assertEqual(ZeroMatrix(self.t1),self.o1)
		self.assertEqual(ZeroMatrix(self.t2),self.o2)


if __name__ == "__main__":


	unittest.main()