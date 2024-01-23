package data

import (
	"math/rand"
)

const adjectiveCnt = 25
const colorCnt = 11
const nounCnt = 13

var adjectives = [adjectiveCnt]string{"pretty", "large", "big", "small", "tall", "short", "long", "handsome",
	"plain", "quaint", "clean", "elegant", "easy", "angry", "crazy", "helpful", "mushy",
	"odd", "unsightly", "adorable", "important", "inexpensive", "cheap", "expensive",
	"fancy"}
var colors = [colorCnt]string{"red", "yellow", "blue", "green", "pink", "brown", "purple", "brown",
	"white", "black", "orange"}
var nouns = [nounCnt]string{"table", "chair", "house", "bbq", "desk", "car", "pony", "cookie",
	"sandwich", "burger", "pizza", "mouse", "keyboard"}

func randAdjective() string {
	return adjectives[rand.Intn(adjectiveCnt)]
}

func randColor() string {
	return colors[rand.Intn(colorCnt)]
}

func randNoun() string {
	return nouns[rand.Intn(nounCnt)]
}

var nextId = 1

type Item struct {
	ID    int
	Label string
}

func BuildData(count int) []*Item {
	data := make([]*Item, count)
	for i := 0; i < count; i++ {
		data[i] = &Item{
			ID:    nextId,
			Label: randAdjective() + " " + randColor() + " " + randNoun(),
		}
		nextId++
	}
	return data
}
