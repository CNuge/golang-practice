import unittest

"""
given two strings, write a method to determine if they are one (or less)
change (insertion, deletion, replacement) away from one another
"""

def OneAway(s1, s2):
	""" take two input strings, return true 
		if they are one change different from one another,
		false if > 1 difference """
	diff = 0

	if len(s1) == len(s2):
		for pos, _ in enumerate(s1):
			if s1[pos] != s2[pos]:
				diff += 1
			if diff > 1:
				return False

	elif (len(s1) - len(s2) == 1) or (len(s1) - len(s2) == -1):
		if len(s1) >  len(s2):
			longer = s1
			shorter = s2
		elif len(s1) < len(s2):
			longer = s2
			shorter = s1
		offset = 0
		for pos, _ in enumerate(shorter):
			if shorter[pos] != longer[pos+offset]:
				diff += 1
				offset += 1
			if diff > 1 :
				return False

	else:
		return False

	return True

class TestOne(unittest.TestCase):
	def test_OneAway(self):
		self.assertTrue(OneAway('timmy', 'jimmy'))
		self.assertTrue(OneAway('dale', 'daleo'))

		self.assertFalse(OneAway('chinook', 'think'))
		self.assertFalse(OneAway('doug', 'dave'))


if __name__ == '__main__':


	unittest.main()