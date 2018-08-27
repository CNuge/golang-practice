import unittest


def IsPalindromePermutation(s):
	""" check if the string s is a permutation of a palindrome"""

	chars = {k : 0 for k in set(s)}

	for i in s:
		chars[i] += 1

	odd_nums = 0

	for k, v in chars.items():
		if (v % 2) != 0:
			odd_nums += 1
		if odd_nums > 1:
			return False

	return True


class TestPalindrome(unittest.TestCase):
	def test_IsPalindromePermutation(self):
		self.assertTrue(IsPalindromePermutation('ceracar'))
		self.assertTrue(IsPalindromePermutation('ahhann'))

		self.assertFalse(IsPalindromePermutation('python'))
		self.assertFalse(IsPalindromePermutation('ceracarit'))



if __name__ == '__main__':


	unittest.main()