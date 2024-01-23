package app

import (
	"main/src/data"
	"strconv"

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
			goui.Component(Jumbotron, JumbotronProps{SetState: setState}).Memo(),
			goui.Element("table", &goui.Attributes{
				Class: "table table-hover table-striped test-data",
				Slot: goui.Element("tbody", &goui.Attributes{
					Slot: goui.Map(state.Items, func(item *data.Item) *goui.Node {
						selected := item.ID == state.Selected
						return goui.
							Component(Row, RowProps{
								Selected: selected,
								Item:     item,
								SetState: setState,
							}).
							Key(strconv.Itoa(item.ID)).
							Memo(selected, item.Label)
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
