package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/cui"
)

const tableData = `OrderDate|Region|Rep|Item|Units|UnitCost|Total
1/6/2017|East|Jones|Pencil|95|1.99|189.05
1/23/2017|Central|Kivell|Binder|50|19.99|999.50
2/9/2017|Central|Jardine|Pencil|36|4.99|179.64
2/26/2017|Central|Gill|Pen|27|19.99|539.73
3/15/2017|West|Sorvino|Pencil|56|2.99|167.44
4/1/2017|East|Jones|Binder|60|4.99|299.40
4/18/2017|Central|Andrews|Pencil|75|1.99|149.25
5/5/2017|Central|Jardine|Pencil|90|4.99|449.10
5/22/2017|West|Thompson|Pencil|32|1.99|63.68
6/8/2017|East|Jones|Binder|60|8.99|539.40
6/25/2017|Central|Morgan|Pencil|90|4.99|449.10
7/12/2017|East|Howard|Binder|29|1.99|57.71
7/29/2017|East|Parent|Binder|81|19.99|1,619.19
8/15/2017|East|Jones|Pencil|35|4.99|174.65
9/1/2017|Central|Smith|Desk|2|125.00|250.00
9/18/2017|East|Jones|Pen Set|16|15.99|255.84
10/5/2017|Central|Morgan|Binder|28|8.99|251.72
10/22/2017|East|Jones|Pen|64|8.99|575.36
11/8/2017|East|Parent|Pen|15|19.99|299.85
11/25/2017|Central|Kivell|Pen Set|96|4.99|479.04
12/12/2017|Central|Smith|Pencil|67|1.29|86.43
12/29/2017|East|Parent|Pen Set|74|15.99|1,183.26
1/15/2018|Central|Gill|Binder|46|8.99|413.54
2/1/2018|Central|Smith|Binder|87|15.00|1,305.00
2/18/2018|East|Jones|Binder|4|4.99|19.96
3/7/2018|West|Sorvino|Binder|7|19.99|139.93
3/24/2018|Central|Jardine|Pen Set|50|4.99|249.50
4/10/2018|Central|Andrews|Pencil|66|1.99|131.34
4/27/2018|East|Howard|Pen|96|4.99|479.04
5/14/2018|Central|Gill|Pencil|53|1.29|68.37
5/31/2018|Central|Gill|Binder|80|8.99|719.20
6/17/2018|Central|Kivell|Desk|5|125.00|625.00
7/4/2018|East|Jones|Pen Set|62|4.99|309.38
7/21/2018|Central|Morgan|Pen Set|55|12.49|686.95
8/7/2018|Central|Kivell|Pen Set|42|23.95|1,005.90
8/24/2018|West|Sorvino|Desk|3|275.00|825.00
9/10/2018|Central|Gill|Pencil|7|1.29|9.03
9/27/2018|West|Sorvino|Pen|76|1.99|151.24
10/14/2018|West|Thompson|Binder|57|19.99|1,139.43
10/31/2018|Central|Andrews|Pencil|14|1.29|18.06
11/17/2018|Central|Jardine|Binder|11|4.99|54.89
12/4/2018|Central|Jardine|Binder|94|19.99|1,879.06
12/21/2018|Central|Andrews|Binder|28|4.99|139.72`

const tableBasic = `[green]func[white] [yellow]main[white]() {
    table := cui.[yellow]NewTable[white]().
        [yellow]SetFixed[white]([red]1[white], [red]1[white])
    [yellow]for[white] row := [red]0[white]; row < [red]40[white]; row++ {
        [yellow]for[white] column := [red]0[white]; column < [red]7[white]; column++ {
            color := tcell.ColorWhite.TrueColor()
            [yellow]if[white] row == [red]0[white] {
                color = tcell.ColorYellow.TrueColor()
            } [yellow]else[white] [yellow]if[white] column == [red]0[white] {
                color = tcell.ColorDarkCyan.TrueColor()
            }
            align := cui.AlignLeft
            [yellow]if[white] row == [red]0[white] {
                align = cui.AlignCenter
            } [yellow]else[white] [yellow]if[white] column == [red]0[white] || column >= [red]4[white] {
                align = cui.AlignRight
            }
            table.[yellow]SetCell[white](row,
                column,
                &cui.TableCell{
                    Text:  [red]"..."[white],
                    Color: color,
                    Align: align,
                })
        }
    }
    cui.[yellow]NewApplication[white]().
        [yellow]SetRoot[white](table, true).
        [yellow]Run[white]()
}`

const tableSeparator = `[green]func[white] [yellow]main[white]() {
    table := cui.[yellow]NewTable[white]().
        [yellow]SetFixed[white]([red]1[white], [red]1[white]).
        [yellow]SetSeparator[white](Borders.Vertical)
    [yellow]for[white] row := [red]0[white]; row < [red]40[white]; row++ {
        [yellow]for[white] column := [red]0[white]; column < [red]7[white]; column++ {
            color := tcell.ColorWhite.TrueColor()
            [yellow]if[white] row == [red]0[white] {
                color = tcell.ColorYellow.TrueColor()
            } [yellow]else[white] [yellow]if[white] column == [red]0[white] {
                color = tcell.ColorDarkCyan.TrueColor()
            }
            align := cui.AlignLeft
            [yellow]if[white] row == [red]0[white] {
                align = cui.AlignCenter
            } [yellow]else[white] [yellow]if[white] column == [red]0[white] || column >= [red]4[white] {
                align = cui.AlignRight
            }
            table.[yellow]SetCell[white](row,
                column,
                &cui.TableCell{
                    Text:  [red]"..."[white],
                    Color: color,
                    Align: align,
                })
        }
    }
    cui.[yellow]NewApplication[white]().
        [yellow]SetRoot[white](table, true).
        [yellow]Run[white]()
}`

const tableBorders = `[green]func[white] [yellow]main[white]() {
    table := cui.[yellow]NewTable[white]().
        [yellow]SetFixed[white]([red]1[white], [red]1[white]).
        [yellow]SetBorders[white](true)
    [yellow]for[white] row := [red]0[white]; row < [red]40[white]; row++ {
        [yellow]for[white] column := [red]0[white]; column < [red]7[white]; column++ {
            color := tcell.ColorWhite.TrueColor()
            [yellow]if[white] row == [red]0[white] {
                color = tcell.ColorYellow.TrueColor()
            } [yellow]else[white] [yellow]if[white] column == [red]0[white] {
                color = tcell.ColorDarkCyan.TrueColor()
            }
            align := cui.AlignLeft
            [yellow]if[white] row == [red]0[white] {
                align = cui.AlignCenter
            } [yellow]else[white] [yellow]if[white] column == [red]0[white] || column >= [red]4[white] {
                align = cui.AlignRight
            }
            table.[yellow]SetCell[white](row,
                column,
                &cui.TableCell{
                    Text:  [red]"..."[white],
                    Color: color,
                    Align: align,
                })
        }
    }
    cui.[yellow]NewApplication[white]().
        [yellow]SetRoot[white](table, true).
        [yellow]Run[white]()
}`

const tableSelectRow = `[green]func[white] [yellow]main[white]() {
    table := cui.[yellow]NewTable[white]().
        [yellow]SetFixed[white]([red]1[white], [red]1[white]).
        [yellow]SetSelectable[white](true, false)
    [yellow]for[white] row := [red]0[white]; row < [red]40[white]; row++ {
        [yellow]for[white] column := [red]0[white]; column < [red]7[white]; column++ {
            color := tcell.ColorWhite.TrueColor()
            [yellow]if[white] row == [red]0[white] {
                color = tcell.ColorYellow.TrueColor()
            } [yellow]else[white] [yellow]if[white] column == [red]0[white] {
                color = tcell.ColorDarkCyan.TrueColor()
            }
            align := cui.AlignLeft
            [yellow]if[white] row == [red]0[white] {
                align = cui.AlignCenter
            } [yellow]else[white] [yellow]if[white] column == [red]0[white] || column >= [red]4[white] {
                align = cui.AlignRight
            }
            table.[yellow]SetCell[white](row,
                column,
                &cui.TableCell{
                    Text:          [red]"..."[white],
                    Color:         color,
                    Align:         align,
                    NotSelectable: row == [red]0[white] || column == [red]0[white],
                })
        }
    }
    cui.[yellow]NewApplication[white]().
        [yellow]SetRoot[white](table, true).
        [yellow]Run[white]()
}`

const tableSelectColumn = `[green]func[white] [yellow]main[white]() {
    table := cui.[yellow]NewTable[white]().
        [yellow]SetFixed[white]([red]1[white], [red]1[white]).
        [yellow]SetSelectable[white](false, true)
    [yellow]for[white] row := [red]0[white]; row < [red]40[white]; row++ {
        [yellow]for[white] column := [red]0[white]; column < [red]7[white]; column++ {
            color := tcell.ColorWhite.TrueColor()
            [yellow]if[white] row == [red]0[white] {
                color = tcell.ColorYellow.TrueColor()
            } [yellow]else[white] [yellow]if[white] column == [red]0[white] {
                color = tcell.ColorDarkCyan.TrueColor()
            }
            align := cui.AlignLeft
            [yellow]if[white] row == [red]0[white] {
                align = cui.AlignCenter
            } [yellow]else[white] [yellow]if[white] column == [red]0[white] || column >= [red]4[white] {
                align = cui.AlignRight
            }
            table.[yellow]SetCell[white](row,
                column,
                &cui.TableCell{
                    Text:          [red]"..."[white],
                    Color:         color,
                    Align:         align,
                    NotSelectable: row == [red]0[white] || column == [red]0[white],
                })
        }
    }
    cui.[yellow]NewApplication[white]().
        [yellow]SetRoot[white](table, true).
        [yellow]Run[white]()
}`

const tableSelectCell = `[green]func[white] [yellow]main[white]() {
    table := cui.[yellow]NewTable[white]().
        [yellow]SetFixed[white]([red]1[white], [red]1[white]).
        [yellow]SetSelectable[white](true, true)
    [yellow]for[white] row := [red]0[white]; row < [red]40[white]; row++ {
        [yellow]for[white] column := [red]0[white]; column < [red]7[white]; column++ {
            color := tcell.ColorWhite.TrueColor()
            [yellow]if[white] row == [red]0[white] {
                color = tcell.ColorYellow.TrueColor()
            } [yellow]else[white] [yellow]if[white] column == [red]0[white] {
                color = tcell.ColorDarkCyan.TrueColor()
            }
            align := cui.AlignLeft
            [yellow]if[white] row == [red]0[white] {
                align = cui.AlignCenter
            } [yellow]else[white] [yellow]if[white] column == [red]0[white] || column >= [red]4[white] {
                align = cui.AlignRight
            }
            table.[yellow]SetCell[white](row,
                column,
                &cui.TableCell{
                    Text:          [red]"..."[white],
                    Color:         color,
                    Align:         align,
                    NotSelectable: row == [red]0[white] || column == [red]0[white],
                })
        }
    }
    cui.[yellow]NewApplication[white]().
        [yellow]SetRoot[white](table, true).
        [yellow]Run[white]()
}`

// Table demonstrates the Table.
func Table(nextSlide func()) (title string, info string, content cui.Primitive) {
	table := cui.NewTable()
	table.SetFixed(1, 1)
	for row, line := range strings.Split(tableData, "\n") {
		for column, cell := range strings.Split(line, "|") {
			color := cui.Styles.PrimaryTextColor
			if row == 0 {
				color = cui.Styles.SecondaryTextColor
			} else if column == 0 {
				color = cui.Styles.TertiaryTextColor
			}
			align := cui.AlignLeft
			if row == 0 {
				align = cui.AlignCenter
			} else if column == 0 || column >= 4 {
				align = cui.AlignRight
			}
			tableCell := cui.NewTableCell(cell)
			tableCell.SetTextColor(color)
			tableCell.SetAlign(align)
			tableCell.SetSelectable(row != 0 && column != 0)
			if column >= 1 && column <= 3 {
				tableCell.SetExpansion(1)
			}
			table.SetCell(row, column, tableCell)
		}
	}
	table.SetBorder(true)
	table.SetTitle("Table")

	code := cui.NewTextView()
	code.SetWrap(false)
	code.SetDynamicColors(true)
	code.SetPadding(1, 1, 2, 0)

	list := cui.NewList()

	basic := func() {
		table.SetBorders(false)
		table.SetSelectable(false, false)
		table.SetSeparator(' ')
		code.Clear()
		fmt.Fprint(code, tableBasic)
	}

	separator := func() {
		table.SetBorders(false)
		table.SetSelectable(false, false)
		table.SetSeparator(cui.Borders.Vertical)
		code.Clear()
		fmt.Fprint(code, tableSeparator)
	}

	borders := func() {
		table.SetBorders(true)
		table.SetSelectable(false, false)
		code.Clear()
		fmt.Fprint(code, tableBorders)
	}

	selectRow := func() {
		table.SetBorders(false)
		table.SetSelectable(true, false)
		table.SetSeparator(' ')
		code.Clear()
		fmt.Fprint(code, tableSelectRow)
	}

	selectColumn := func() {
		table.SetBorders(false)
		table.SetSelectable(false, true)
		table.SetSeparator(' ')
		code.Clear()
		fmt.Fprint(code, tableSelectColumn)
	}

	selectCell := func() {
		table.SetBorders(false)
		table.SetSelectable(true, true)
		table.SetSeparator(' ')
		code.Clear()
		fmt.Fprint(code, tableSelectCell)
	}

	navigate := func() {
		app.SetFocus(table)
		table.SetDoneFunc(func(key tcell.Key) {
			app.SetFocus(list)
		})
		table.SetSelectedFunc(func(row int, column int) {
			app.SetFocus(list)
		})
	}

	list.ShowSecondaryText(false)
	list.SetPadding(1, 1, 2, 2)

	var demoTableText = []struct {
		text     string
		shortcut rune
		selected func()
	}{
		{"Basic table", 'b', basic},
		{"Table with separator", 's', separator},
		{"Table with borders", 'o', borders},
		{"Selectable rows", 'r', selectRow},
		{"Selectable columns", 'c', selectColumn},
		{"Selectable cells", 'l', selectCell},
		{"Navigate", 'n', navigate},
		{"Next slide", 'x', nextSlide},
	}

	for _, tableText := range demoTableText {
		item := cui.NewListItem(tableText.text)
		item.SetShortcut(tableText.shortcut)
		item.SetSelectedFunc(tableText.selected)
		list.AddItem(item)
	}

	basic()

	subFlex := cui.NewFlex()
	subFlex.SetDirection(cui.FlexRow)
	subFlex.AddItem(list, 10, 1, true)
	subFlex.AddItem(table, 0, 1, false)

	flex := cui.NewFlex()
	flex.AddItem(subFlex, 0, 1, true)
	flex.AddItem(code, codeWidth, 1, false)

	return "Table", "", flex
}
