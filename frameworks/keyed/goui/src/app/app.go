package app

import (
	"main/src/data"

	"github.com/twharmon/goui"
)

type AppState struct {
	Items    []*data.Item
	Selected int
}

func App(goui.NoProps) *goui.Node {
	state, setState := goui.UseState(&AppState{})

	return goui.Element("div", &goui.Attributes{
		Class: "container",
		Slot: []*goui.Node{
			goui.Component(Jumbotron, JumbotronProps{SetState: setState}),
			goui.Element("table", &goui.Attributes{
				Class: "table table-hover table-striped test-data",
				Slot: goui.Element("tbody", &goui.Attributes{
					Slot: goui.Map(state.Items, func(item *data.Item) *goui.Node {
						return goui.Component(Row, RowProps{
							Selected: item.ID == state.Selected,
							Item:     item,
							SetState: setState,
						})
					}),
				}),
			}),
			preloadIcon,
		},
	})
}

var preloadIcon = goui.Element("span", &goui.Attributes{
	Class:      "preloadicon glyphicon glyphicon-remove",
	AriaHidden: true,
})
