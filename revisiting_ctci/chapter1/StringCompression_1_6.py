"""
take a string and compress it into repeated digit counts:

AAAATTTTTGGG
A4T5G3

if the new string is longer than the original, then return the original.

"""

import unittest


def StringCompression(s):
	outstring = ''
	current_char = ''
	current_count = 0

	for i in s:
		if i == current_char:
			current_count += 1

		else:
			if current_count != 0:
				outstring+= current_char
				outstring+=str(current_count)

			current_char = i
			current_count = 1

	outstring+= current_char
	outstring+=str(current_count)

	if len(outstring) < len(s):
		return outstring
	return s


class TestStringCompression(unittest.TestCase):
	def test_StringCompression(self):
		t1 = "AAAATTTTTGGG"
		o1 = "A4T5G3"
		t2 = "ATGCATGCATGC"
		o2 = "ATGCATGCATGC"
		t3 = "CGGGAAAATTTTTGGG"
		o3 = "C1G3A4T5G3"

		self.assertEqual(StringCompression(t1), o1)
		self.assertEqual(StringCompression(t2), o2)
		self.assertEqual(StringCompression(t3), o3)

if __name__ == '__main__':
	unittest.main()