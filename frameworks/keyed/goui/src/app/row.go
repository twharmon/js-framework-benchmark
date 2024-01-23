package app

import (
	"main/src/data"

	"github.com/twharmon/goui"
)

type RowProps struct {
	Selected bool
	SetState func(func(*AppState) *AppState)
	Item     *data.Item
}

func Row(props RowProps) *goui.Node {
	id := props.Item.ID

	handleSelect := goui.UseCallback(func(*goui.MouseEvent) {
		props.SetState(func(state *AppState) *AppState {
			return &AppState{
				Selected: id,
				Items:    state.Items,
			}
		})
	}, goui.Deps{id})

	handleDelete := goui.UseCallback(func(*goui.MouseEvent) {
		props.SetState(func(state *AppState) *AppState {
			for i := range state.Items {
				if state.Items[i].ID == id {
					state.Items = append(state.Items[:i], state.Items[i+1:]...)
					return &AppState{
						Selected: state.Selected,
						Items:    state.Items,
					}
				}
			}
			return state
		})
	}, goui.Deps{id})

	trClass := ""
	if props.Selected {
		trClass = "danger"
	}

	return goui.Element("tr", &goui.Attributes{
		Class: trClass,
		Slot: []*goui.Node{
			goui.Element("td", &goui.Attributes{
				Class: "col-md-1",
				Slot:  id,
			}),
			goui.Element("td", &goui.Attributes{
				Class: "col-md-4",
				Slot: goui.Element("a", &goui.Attributes{
					OnClick: handleSelect,
					Slot:    props.Item.Label,
				}),
			}),
			goui.Element("td", &goui.Attributes{
				Class: "col-md-1",
				Slot: goui.Element("a", &goui.Attributes{
					OnClick: handleDelete,
					Slot:    trashIcon,
				}),
			}),
			emptyTd,
		},
	})
}

var emptyTd = goui.Element("td", &goui.Attributes{Class: "col-md-6"})
var trashIcon = goui.Element("span", &goui.Attributes{
	Class:      "glyphicon glyphicon-remove",
	AriaHidden: true,
})
