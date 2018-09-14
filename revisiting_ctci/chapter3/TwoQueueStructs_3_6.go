/*

animal shelter holds dogs and cats.

first in shelter must be first adopted (queue)

people can select whether they want a dog or a cat,
otherwise if no choice they get the oldest.


implement this setup

*/

package main


import (
	"fmt"
)


type Queue struct {
	ord []Animal
}

// representation of the Queue, just the order
func (q Queue) String() string {
	return fmt.Sprintf("%v", q.ord)
}

// add an integer to the back of the list
func (q *Queue) Push(item Animal)  {
	q.ord = append(q.ord, item)
}

// remove
func (q *Queue) Pop() Animal {
	output := q.ord[0]
	q.ord = q.ord[1:]
	return output
}

// look at the first position in the Queue
func (q *Queue) Peek() Animal {
	return q.ord[0]
}


type Animal struct {
	Species string
	Id		int
}

type Shelter struct {
	Dogs Queue
	Cats Queue
	Id	int
}

// take a new animal instance and add it in to the shelter
func (s Shelter) AddAnimal(animal_type string){
	s.Id++
	new_animal := Animal{Species: animal_type ,
							Id: s.Id }
	if new_animal.Species == "Cat" {
		s.Cats.Push(new_animal)
	} else if new_animal.Species == "Dog" {
		s.Dogs.Push(new_animal)
	}
}

func (s Shelter) Adopt(species string) Animal {
	
	if species == "Dog"{
		return s.Dogs.Pop()
	}else if species == "Cat"{
		return s.Cats.Pop()
	}else{
		sr_cat := s.Cats.Peek() // peek at the oldest cat in the queue
		sr_dog := s.Dogs.Peek() // peek at the oldest dog in the queue
		
		if sr_cat.Id < sr_dog.Id {
			return s.Cats.Pop()
		} else if  sr_cat.Id > sr_dog.Id{
			return s.Dogs.Pop()
		} 
	}
	return Animal{}
}


func main(){

	test_shelter := Shelter{}

	test_shelter.AddAnimal("Dog")

	test_shelter.AddAnimal("Cat")
	test_shelter.AddAnimal("Cat")
	test_shelter.AddAnimal("Cat")
	test_shelter.AddAnimal("Cat")

	test_shelter.AddAnimal("Dog")
	test_shelter.AddAnimal("Dog")
	test_shelter.AddAnimal("Dog")
	test_shelter.AddAnimal("Dog")

	out1 := test_shelter.Adopt("any")
	out2 := test_shelter.Adopt("any")
	out3 := test_shelter.Adopt("any")

	fmt.Printf("%v == Dog\n",out1)
	fmt.Printf("%v == Cat\n",out2)
	fmt.Printf("%v == Cat\n",out3)

}