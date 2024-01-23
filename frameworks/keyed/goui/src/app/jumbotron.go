package app

import (
	"main/src/data"

	"github.com/twharmon/goui"
)

type JumbotronProps struct {
	SetState func(func(*AppState) *AppState)
}

func button(id string, txt string, onclick *goui.Callback[func(*goui.MouseEvent)]) *goui.Node {
	return goui.Element("div", &goui.Attributes{
		Class: "col-sm-6 smallpad",
		Slot: goui.Element("button", &goui.Attributes{
			ID:      id,
			OnClick: onclick,
			Type:    "button",
			Class:   "btn btn-primary btn-block",
			Slot:    txt,
		}),
	})
}

func Jumbotron(props JumbotronProps) *goui.Node {
	handleCreate1k := goui.UseCallback(func(*goui.MouseEvent) {
		props.SetState(func(state *AppState) *AppState {
			return &AppState{
				Selected: 0,
				Items:    data.BuildData(1000),
			}
		})
	}, goui.Deps{})

	handleCreate10k := goui.UseCallback(func(*goui.MouseEvent) {
		props.SetState(func(state *AppState) *AppState {
			return &AppState{
				Selected: 0,
				Items:    data.BuildData(10000),
			}
		})
	}, goui.Deps{})

	handleAppend1k := goui.UseCallback(func(*goui.MouseEvent) {
		props.SetState(func(state *AppState) *AppState {
			return &AppState{
				Selected: state.Selected,
				Items:    append(state.Items, data.BuildData(1000)...),
			}
		})
	}, goui.Deps{})

	handleUpdateEvery10th := goui.UseCallback(func(*goui.MouseEvent) {
		props.SetState(func(state *AppState) *AppState {
			for i := 0; i < len(state.Items); i += 10 {
				state.Items[i].Label += " !!!"
			}
			return &AppState{
				Selected: state.Selected,
				Items:    state.Items,
			}
		})
	}, goui.Deps{})

	handleClear := goui.UseCallback(func(*goui.MouseEvent) {
		props.SetState(func(*AppState) *AppState {
			return &AppState{}
		})
	}, goui.Deps{})

	handleSwapRows := goui.UseCallback(func(*goui.MouseEvent) {
		props.SetState(func(state *AppState) *AppState {
			if len(state.Items) >= 999 {
				state.Items[1], state.Items[998] = state.Items[998], state.Items[1]
			}
			return &AppState{
				Selected: state.Selected,
				Items:    state.Items,
			}
		})
	}, goui.Deps{})

	return goui.Element("div", &goui.Attributes{
		Class: "jumbotron",
		Slot: []*goui.Node{
			goui.Element("div", &goui.Attributes{
				Class: "row",
				Slot: []*goui.Node{
					heading,
					goui.Element("div", &goui.Attributes{
						Class: "col-md-6",
						Slot: goui.Element("div", &goui.Attributes{
							Class: "row",
							Slot: []*goui.Node{
								button("run", "Create 1,000 rows", handleCreate1k),
								button("runlots", "Create 10,000 rows", handleCreate10k),
								button("add", "Append 1,000 rows", handleAppend1k),
								button("update", "Update every 10th row", handleUpdateEvery10th),
								button("clear", "Clear", handleClear),
								button("swaprows", "Swap rows", handleSwapRows),
							},
						}),
					}),
				},
			}),
		},
	})
}

var heading = goui.Element("div", &goui.Attributes{
	Class: "col-md-6",
	Slot:  goui.Element("h1", &goui.Attributes{Slot: "GoUI"}),
})
