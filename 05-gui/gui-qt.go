package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	// Vytvoření nové Qt aplikace
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Vytvoření hlavního okna
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Go Qt Example")
	window.SetMinimumSize2(400, 300)

	// Vytvoření centrálního widgetu a layoutu
	centralWidget := widgets.NewQWidget(nil, 0)
	layout := widgets.NewQVBoxLayout()

	// Vytvoření tlačítka
	button := widgets.NewQPushButton2("Click Me!", nil)
	button.ConnectClicked(func(bool) {
		widgets.QMessageBox_Information(nil, "Message", "Button clicked!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	// Přidání tlačítka do layoutu
	layout.AddWidget(button, 0, 0)

	// Nastavení layoutu do centrálního widgetu
	centralWidget.SetLayout(layout)

	// Nastavení centrálního widgetu do hlavního okna
	window.SetCentralWidget(centralWidget)

	// Zobrazení okna
	window.Show()

	// Spuštění aplikace
	app.Exec()
}

