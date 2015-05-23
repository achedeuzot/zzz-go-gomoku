package gui

import ()

type Scene interface {
	PlayScene()
	handleEvents()
}

var CurrScene Scene
var SceneMap map[string]Scene = make(map[string]Scene)

func initScenes() {
	SceneMap["MenuMain"] = NewMainMenu()
	CurrScene = SceneMap["MenuMain"]
	// XXX temporary
	SceneMap["Game"] = NewBoard()
	SceneMap["Options"] = NewOptions()
	// CurrScene = SceneMap["Board"]
}
