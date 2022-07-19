package main

import (
	"fmt"

	"merge_lists/word_wrap"
)

func main() {
	width := 50
	text := "The prize was delivered to Tom with as much effusion as the superintendent could pump up under under the circumstances; but it lacked somewhat of the true gush, for the poor fellow’s instinct taught him that there was a mystery here that could not well bear the light, perhaps; it was simply preposterous that this boy had warehoused two thousand sheaves of Scriptural wisdom on his premises—a dozen would strain his capacity, without a doubt.\n\nAmy Lawrence was proud and glad, and she tried to make Tom see it in her face—but he wouldn’t look. She wondered; then she was just a grain troubled; next a dim suspicion came and went—came again; she watched; a furtive glance told her worlds—and then her heart broke, and she was jealous, and angry, and the tears came and she hated everybody. Tom most of all (she thought)."
	result := word_wrap.WordWrap(text, width)

	fmt.Printf(result)
}
