"""

animal shelter holds dogs and cats.

first in shelter must be first adopted (queue)

people can select whether they want a dog or a cat,
otherwise if no choice they get the oldest.


implement this setup


"""

from queue import Queue


class Shelter():
	def __init__(self):
		self.Cats = Queue()
		self.Dogs = Queue()
		self.id = 0

	def Add(self, animal, species):
		"""take a cat or dog and its name, add to the shelter"""
		self.id +=1
		if species == 'Cat':
			self.Cats.Push((animal, species , self.id))  
		elif species == 'Dog':
			self.Dogs.Push((animal, species , self.id))
		else:
			raise ValueError('Must specify the species of the animal. String: Cat or Dog')

	def Adopt(self, preference = None):
		""" if preference given, adopt oldest member of that queue, 
			otherwise compare data of the two and adopt lower id number """
		if preference == 'Cat':
			return self.Cats.Pop()
		elif preference == 'Dog':
			return self.Dogs.Pop()
		elif preference == None:
			cat_last = self.Cats.Peek()
			dog_last = self.Dogs.Peek()
			if cat_last[-1] < dog_last [-1]:
				return self.Cats.Pop()
			else:
				return self.Dogs.Pop()

		else:
			raise ValueError('Must specify the preferred species of the animal. String: Cat or Dog\n'+\
								'If no preference then no value should be passed.')


if __name__ == '__main__':

	test_shelter = Shelter()

	test_shelter.Add('Lassie', 'Dog')

	test_shelter.Add('Garfield', 'Cat')
	test_shelter.Add('Andrew', 'Cat')
	test_shelter.Add('Lloyd', 'Cat')
	test_shelter.Add('Weber', 'Cat')

	test_shelter.Add('Mr. Pickles', 'Dog')

	test_shelter.Add('Alfie', 'Dog')

	print(test_shelter.Adopt('Cat'))
	print(test_shelter.Adopt())
	print(test_shelter.Adopt('Cat'))
	print(test_shelter.Adopt('Dog'))
	print(test_shelter.Adopt())
	print(test_shelter.Adopt('Dog'))

