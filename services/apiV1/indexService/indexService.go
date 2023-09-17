package indexservice

import "math/rand"

type IndexService struct {
	welcomeQoutes []string
}

func NewIndexService() *IndexService {
	return &IndexService{
		welcomeQoutes: []string{
			"Hello, World!",
			"Hey, baby! Anybody ever tells you I have beautiful eyes?",
			"But enough about me. Let's talk about me.",
			"I am Johnny Bravo, the one-man army!",
			"Mama Mia. That's a spicy meatball!",
			"Now remember, I do my best work when I'm being worshiped as a god.",
			"I'm so pretty I could spend hours just looking at myself in the mirror.",
			"I'm the most popular guy in town. All the ladies love me.",
			"I'm so smooth I could talk a snowman into a snowball fight.",
			"I'm so cool I could freeze the Sahara Desert.",
			"I'm so strong I could lift a building with my bare hands.",
			"I'm so smart I could solve the Rubik's Cube in my sleep.",
			"I'm so funny I could make a grown man cry.",
			"I'm so awesome I could make a unicorn jealous.",
		},
	}
}

func (is *IndexService) RandomWelcomeMessage() string {
	return is.welcomeQoutes[rand.Intn(len(is.welcomeQoutes))]
}
