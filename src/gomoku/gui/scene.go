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
	SceneMap["Game"] = NewGame()
	SceneMap["Options"] = NewOptions()
	CurrScene = SceneMap["MenuMain"]
}
