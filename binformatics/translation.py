"""
read in the an rna seqenece,
find the reference frame by searching for a start codon 
and transcribe it until a stop codon is reached
"""

trans_map = {"UUU":"F", "UUC":"F", "UUA":"L", "UUG":"L",
				"UCU":"S", "UCC":"S", "UCA":"S", "UCG":"S",
				"UAU":"Y", "UAC":"Y", "UAA":"STOP", "UAG":"STOP",
				"UGU":"C", "UGC":"C", "UGA":"STOP", "UGG":"W",
				"CUU":"L", "CUC":"L", "CUA":"L", "CUG":"L",
				"CCU":"P", "CCC":"P", "CCA":"P", "CCG":"P",
				"CAU":"H", "CAC":"H", "CAA":"Q", "CAG":"Q",
				"CGU":"R", "CGC":"R", "CGA":"R", "CGG":"R",
				"AUU":"I", "AUC":"I", "AUA":"I", "AUG":"M",
				"ACU":"T", "ACC":"T", "ACA":"T", "ACG":"T",
				"AAU":"N", "AAC":"N", "AAA":"K", "AAG":"K",
				"AGU":"S", "AGC":"S", "AGA":"R", "AGG":"R",
				"GUU":"V", "GUC":"V", "GUA":"V", "GUG":"V",
				"GCU":"A", "GCC":"A", "GCA":"A", "GCG":"A",
				"GAU":"D", "GAC":"D", "GAA":"E", "GAG":"E",
				"GGU":"G", "GGC":"G", "GGA":"G", "GGG":"G",}


class Transcript():
	def __init__(self, rna):
		""" pass in an rna string, it will be translated and 
			stored as Transcript.protein """
		self.rna = rna
		self.protein = ""
		self.translate()

	def __repr__(self):
		return f"rna: {self.rna}\nprotein: {self.protein}\n"

	def frame(self):
		""" find the first bp following the start codon """
		pos = 0
		while pos <= len(self.rna):
			if self.rna[pos:pos+3] == "AUG":
				return  pos + 3
			pos += 1
		return pos

	def translate(self):
		""" take the passed in rna and determine the protein seq """
		pos = self.frame()
		while pos <= len(self.rna) - 3:
			aa = trans_map[self.rna[pos: pos + 3]]
			if aa == 'STOP':
				break
			else:
				self.protein += aa
				pos += 3
		else:
			""" this will be triggered only if the while loop doesn't
				encounter a break, i.e. a stop codon """
			self.protein += '...'

if __name__ == "__main__":
	transcript1 = Transcript("AAAUAUGGCCGCAGAGUAGGGGGG" ) # AAE
	print(transcript1.protein)
	transcript2 = Transcript("AAUGUUUUUUUUUUUUUUUUUUUU") # no stop FFFFF...
	print(transcript2.protein)
	transcript3 = Transcript("AAUGUUUAAAUUUAAAGAAUAG" ) # FKFKE
	print(transcript3.protein)
	transcript4 = Transcript("UUAAAUUUAAAGAAUAG") # no start
	print(transcript4.protein)
	print(transcript1)
