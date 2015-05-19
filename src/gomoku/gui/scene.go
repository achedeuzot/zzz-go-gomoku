package gui

import ()

type Scene interface {
	PlayScene()
}

var CurrScene Scene
var SceneMap map[string]Scene = make(map[string]Scene)

func initScenes() {
	SceneMap["MenuMain"] = NewMainMenu()
	CurrScene = SceneMap["MenuMain"]
	// XXX temporary
	SceneMap["Board"] = NewBoard()
	CurrScene = SceneMap["Board"]
}
