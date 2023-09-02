package puppy

import "github.com/vladfreishmidt/dog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From11() string {
	return "Hey, I'm from version 1.1.0"
}

func From12() string {
	return "Hey, I'm from version 1.2.0"
}
