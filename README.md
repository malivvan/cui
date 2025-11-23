# cview - Terminal-based user interface toolkit
[![GoDoc](https://codeberg.org/tslocum/godoc-static/raw/branch/master/badge.svg)](https://docs.rocket9labs.com/github.com/malivvan/cui)
[![Donate](https://img.shields.io/liberapay/receives/rocket9labs.com.svg?logo=liberapay)](https://liberapay.com/rocket9labs.com)

This package is a fork of [tview](https://github.com/rivo/tview).
See [FORK.md](https://github.com/malivvan/cui/src/branch/master/FORK.md) for more information.

## Demo

`ssh cui.rocket9labs.com -p 20000`

[![Recording of presentation demo](https://github.com/malivvan/cui/raw/branch/master/cui.svg)](https://github.com/malivvan/cui/src/branch/master/demos/presentation)

## Widgets

This section provides comprehensive documentation of all available widgets in the cui package.

### Core Widgets

#### Box
The base primitive for all widgets. Provides background color, borders, titles, and padding. All other widgets embed Box and inherit its functionality.

**Key Features:**
- Configurable borders and titles
- Padding and margin control
- Background colors and transparency
- Focus management

**Demo:** [`_demo/box`](https://github.com/malivvan/cui/tree/master/_demo/box)

#### Application
The top-level container that manages the event loop, screen drawing, and focus handling for your terminal application.

**Key Features:**
- Event loop management
- Mouse and keyboard input handling
- Screen updates and rendering
- Focus management across widgets

### Input Widgets

#### Button
A labeled button that triggers an action when selected or clicked.

**Key Features:**
- Customizable label and styling
- Keyboard and mouse activation
- Focus indicators
- Selected callback handlers

**Demo:** [`_demo/button`](https://github.com/malivvan/cui/tree/master/_demo/button)

#### CheckBox
A selectable checkbox for boolean values with optional label and message text.

**Key Features:**
- Toggle on/off states
- Customizable labels
- Keyboard shortcuts
- Change event handlers

**Demo:** [`_demo/checkbox`](https://github.com/malivvan/cui/tree/master/_demo/checkbox)

#### InputField
A single-line text entry field with support for masked input (passwords), autocomplete, and validation.

**Key Features:**
- Text input and editing
- Password masking
- Autocomplete support
- Input validation
- Field width configuration

**Demos:** 
- [`_demo/inputfield/simple`](https://github.com/malivvan/cui/tree/master/_demo/inputfield/simple)
- [`_demo/inputfield/autocomplete`](https://github.com/malivvan/cui/tree/master/_demo/inputfield/autocomplete)
- [`_demo/inputfield/autocompleteasync`](https://github.com/malivvan/cui/tree/master/_demo/inputfield/autocompleteasync)

#### DropDown
A drop-down selection field allowing users to choose from a list of options.

**Key Features:**
- Multiple option support
- Current selection display
- Keyboard navigation
- Custom option callbacks
- Optional labels

**Demo:** [`_demo/dropdown`](https://github.com/malivvan/cui/tree/master/_demo/dropdown)

#### Slider
An interactive progress bar that can be modified via keyboard and mouse input.

**Key Features:**
- Adjustable value via keyboard/mouse
- Configurable range
- Optional label
- Horizontal and vertical orientation
- Custom styling

#### Form
A container for form elements including input fields, dropdowns, checkboxes, and buttons.

**Key Features:**
- Multiple field types support
- Automatic layout
- Tab navigation between fields
- Form submission handling
- Cancel/submit buttons

**Demo:** [`_demo/form`](https://github.com/malivvan/cui/tree/master/_demo/form)

### Display Widgets

#### TextView
A scrollable widget for displaying multi-colored, formatted text with optional highlighting and regions.

**Key Features:**
- Multi-color text support
- Text wrapping and word wrapping
- Scrolling support
- Highlighting and regions
- Dynamic content updates
- io.Writer interface support

**Demo:** [`_demo/textview`](https://github.com/malivvan/cui/tree/master/_demo/textview)

#### List
A navigable list of items with optional keyboard shortcuts and context menus.

**Key Features:**
- Main and secondary text per item
- Keyboard shortcuts
- Item selection callbacks
- Context menus
- Scrollable content
- Search/filter support

**Demo:** [`_demo/list`](https://github.com/malivvan/cui/tree/master/_demo/list)

#### Table
A scrollable table widget displaying tabular data with support for cell highlighting, fixed rows/columns, and sorting.

**Key Features:**
- Fixed header rows and columns
- Cell selection and navigation
- Custom cell styling
- Expandable cells
- Sorting support
- Scrollbars

**Demo:** [`_demo/table`](https://github.com/malivvan/cui/tree/master/_demo/table)

#### TreeView
A hierarchical tree widget for displaying and navigating tree structures.

**Key Features:**
- Expandable/collapsible nodes
- Node selection and navigation
- Custom node colors and text
- Reference data per node
- Keyboard navigation

**Demo:** [`_demo/treeview`](https://github.com/malivvan/cui/tree/master/_demo/treeview)

#### Image
A widget that displays images in the terminal using graphical characters.

**Key Features:**
- Image rendering with dithering
- Size and color configuration
- Aspect ratio preservation
- Multiple image formats support

**Demo:** [`_demo/image`](https://github.com/malivvan/cui/tree/master/_demo/image)

#### ProgressBar
A widget indicating progress of an operation with horizontal or vertical orientation.

**Key Features:**
- Progress value display
- Horizontal and vertical modes
- Custom fill and empty characters
- Color customization

**Demo:** [`_demo/progressbar`](https://github.com/malivvan/cui/tree/master/_demo/progressbar)

### Layout Widgets

#### Flex
A Flexbox-based layout container arranging primitives horizontally or vertically with proportional or fixed sizing.

**Key Features:**
- Row or column layout
- Proportional and fixed sizing
- Nested flex layouts
- Dynamic resizing

**Demo:** [`_demo/flex`](https://github.com/malivvan/cui/tree/master/_demo/flex)

#### Grid
A grid-based layout manager for positioning widgets in rows and columns.

**Key Features:**
- Row and column sizing
- Widget spanning multiple cells
- Minimum size constraints
- Responsive layouts
- Scrollable content

**Demo:** [`_demo/grid`](https://github.com/malivvan/cui/tree/master/_demo/grid)

#### Panels
A container for managing multiple primitives with visibility control and layering.

**Key Features:**
- Multiple panel support
- Show/hide panels
- Panel stacking (z-order)
- Easy switching between panels

**Demo:** [`_demo/panels`](https://github.com/malivvan/cui/tree/master/_demo/panels)

#### TabbedPanels
A tabbed container with tab navigation for switching between panels.

**Key Features:**
- Multiple tabs with labels
- Tab switching via keyboard/mouse
- Vertical or horizontal tab placement
- Tab customization

**Demo:** [`_demo/tabbedpanels`](https://github.com/malivvan/cui/tree/master/_demo/tabbedpanels)

#### Frame
A wrapper widget that adds space, headers, and footers around another primitive.

**Key Features:**
- Customizable borders
- Header and footer text
- Spacing control
- Nested primitive support

**Demo:** [`_demo/frame`](https://github.com/malivvan/cui/tree/master/_demo/frame)

### Window Management

#### Window
A draggable and resizable frame around a primitive, meant to be used with WindowManager.

**Key Features:**
- Drag and drop support
- Resizable borders
- Fullscreen toggle
- Focus management

**Demo:** [`_demo/window`](https://github.com/malivvan/cui/tree/master/_demo/window)

#### WindowManager
A container for managing multiple windows with focus control and window ordering.

**Key Features:**
- Multiple window management
- Window focus control
- Z-order management
- Window stacking

**Demo:** [`_demo/window`](https://github.com/malivvan/cui/tree/master/_demo/window)

### Dialog Widgets

#### Modal
A centered modal dialog for displaying messages and prompting for user decisions.

**Key Features:**
- Centered display
- Customizable buttons
- Message text
- Block interaction with other widgets

**Demo:** [`_demo/modal`](https://github.com/malivvan/cui/tree/master/_demo/modal)

#### ContextMenu
A menu that appears on user interaction (e.g., right-click) providing context-specific actions.

**Key Features:**
- Dynamic positioning
- Custom menu items
- Keyboard and mouse navigation
- Selection callbacks

### Chart Widgets

The `chart` package provides specialized widgets for data visualization.

#### BarChart
A bar chart widget for displaying categorical data with vertical bars.

**Key Features:**
- Multiple bar support
- Custom colors per bar
- Axis labels
- Auto-scaling

**Demo:** [`_demo/chart_barchart`](https://github.com/malivvan/cui/tree/master/_demo/chart_barchart)

#### Sparkline
A compact line chart widget for showing trends in small spaces.

**Key Features:**
- Compact data visualization
- Line color customization
- Data title support
- Auto-scaling

**Demo:** [`_demo/chart_sparkline`](https://github.com/malivvan/cui/tree/master/_demo/chart_sparkline)

#### Plot
A plotting widget supporting line charts and scatter plots with axis labels.

**Key Features:**
- Line chart and scatter plot modes
- X and Y axis labels
- Multiple data series
- Custom markers (braille or dot)
- Axis range configuration

**Demos:**
- [`_demo/chart_plot`](https://github.com/malivvan/cui/tree/master/_demo/chart_plot)
- [`_demo/chart_plot_custom_range`](https://github.com/malivvan/cui/tree/master/_demo/chart_plot_custom_range)
- [`_demo/chart_plot_xaxis_labels`](https://github.com/malivvan/cui/tree/master/_demo/chart_plot_xaxis_labels)

#### ActivityModeGauge
A gauge widget displaying animated activity progress.

**Key Features:**
- Animated progress indicator
- Customizable colors
- Activity tracking

**Demo:** [`_demo/chart_gauge_am`](https://github.com/malivvan/cui/tree/master/_demo/chart_gauge_am)

#### PercentageModeGauge
A gauge displaying percentage-based progress with visual feedback.

**Key Features:**
- Percentage display
- Current value tracking
- Color customization

**Demo:** [`_demo/chart_gauge_pm`](https://github.com/malivvan/cui/tree/master/_demo/chart_gauge_pm)

#### UtilModeGauge
A gauge for displaying utilization metrics with warning and critical thresholds.

**Key Features:**
- Threshold-based coloring (ok/warning/critical)
- Label support
- Percentage display
- Configurable thresholds

**Demo:** [`_demo/chart_gauge_um`](https://github.com/malivvan/cui/tree/master/_demo/chart_gauge_um)

#### Spinner
An animated spinner widget for indicating loading or processing states.

**Key Features:**
- Multiple spinner styles
- Animation support
- Style customization

**Demo:** [`_demo/chart_spinner`](https://github.com/malivvan/cui/tree/master/_demo/chart_spinner)

#### MessageDialog
A dialog widget for displaying information or error messages with buttons.

**Key Features:**
- Info and error dialog types
- Custom messages
- Button callbacks
- Centered display

**Demo:** [`_demo/chart_dialog`](https://github.com/malivvan/cui/tree/master/_demo/chart_dialog)

### Editor Widget

The `editor` package provides a full-featured text editor widget.

#### View (Editor)
A powerful text editor widget with syntax highlighting, line numbers, and editing capabilities.

**Key Features:**
- Text editing with cursor support
- Syntax highlighting
- Line numbers
- Multiple buffers
- Undo/redo support
- Search functionality

**Demo:** [`_demo/editor`](https://github.com/malivvan/cui/tree/master/_demo/editor)

### Terminal Emulator

The `vte` package provides a virtual terminal emulator widget.

#### VT (Virtual Terminal)
A terminal emulator widget that can run shell commands and display their output.

**Key Features:**
- Full terminal emulation
- ANSI escape sequence support
- Shell command execution
- Scrollback buffer
- Mouse support
- Sixel graphics support

**Demos:**
- [`_demo/vte`](https://github.com/malivvan/cui/tree/master/_demo/vte)
- [`_demo/vte_parse`](https://github.com/malivvan/cui/tree/master/_demo/vte_parse)
- [`_demo/vte_show_keys`](https://github.com/malivvan/cui/tree/master/_demo/vte_show_keys)

### Menu Widget

The `menu` package provides menu bar functionality.

#### Menu/MenuItem
A menu bar widget with hierarchical menu items and submenus.

**Key Features:**
- Menu bar support
- Nested submenus
- Menu item actions
- Keyboard navigation

## Features

All widgets may be customized and extended to suit any application.

[Mouse support](https://docs.rocket9labs.com/github.com/malivvan/cui#hdr-Mouse_Support) is available.

## Applications

A list of applications powered by cview is available via [pkg.go.dev](https://pkg.go.dev/github.com/malivvan/cui?tab=importedby).

## Installation

```bash
go get github.com/malivvan/cui
```

## Hello World

This basic example creates a TextView titled "Hello, World!" and displays it in your terminal:

```go
package main

import (
	"github.com/malivvan/cui"
)

func main() {
	app := cui.NewApplication()

	tv := cui.NewTextView()
	tv.SetBorder(true)
	tv.SetTitle("Hello, world!")
	tv.SetText("Lorem ipsum dolor sit amet")
	
	app.SetRoot(tv, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
```

Examples are available via [godoc](https://docs.rocket9labs.com/github.com/malivvan/cui#pkg-examples)
and in the [demos](https://github.com/malivvan/cui/src/branch/master/demos) directory.

For a presentation highlighting the features of this package, compile and run
the program in the [demos/presentation](https://github.com/malivvan/cui/src/branch/master/demos/presentation) directory.

## Documentation

Package documentation is available via [godoc](https://docs.rocket9labs.com/github.com/malivvan/cui).

An [introduction tutorial](https://rocket9labs.com/post/tview-and-you/) is also available.

## Dependencies

This package is based on [github.com/gdamore/tcell](https://github.com/gdamore/tcell)
(and its dependencies) and [github.com/rivo/uniseg](https://github.com/rivo/uniseg).

## Support

[CONTRIBUTING.md](https://github.com/malivvan/cui/src/branch/master/CONTRIBUTING.md) describes how to share
issues, suggestions and patches (pull requests).

## Packages
- / [codeberg.org/tslocum/cview](https://codeberg.org/tslocum/cview/src/commit/242e7c1f1b61a4b3722a1afb45ca1165aefa9a59)
- /bind.go [codeberg.org/tslocum/cbind](https://codeberg.org/tslocum/cbind/src/commit/5cd49d3cfccbe4eefaab8a5282826aa95100aa42)
- /vte/ [git.sr.ht/~rockorager/tcell-term](https://git.sr.ht/~rockorager/tcell-term/refs/v0.10.0)
- /femto/ [github.com/wellcomez/femto](https://github.com/wellcomez/femto/tree/8413a0288bcb042fd0de5cbbcb9893c16a01ee69)
- /chart/ [github.com/navidys/tvxwidgets](https://github.com/navidys/tvxwidgets/commit/96bcc0450684693eebd4f8e3e95fcc40eae2dbaa)
