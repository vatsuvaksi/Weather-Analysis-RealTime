package network

type ResponseQoutes struct {
	WelcomeMessage string `json:"welcomeMessage"`
	WordsOfWisdom  string `json:"wordsOfWisdom"`
}
type FailiureResponse struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}
