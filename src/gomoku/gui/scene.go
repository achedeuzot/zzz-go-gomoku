package gui

import ()

type Scene interface {
	PlayScene()
}

var CurrScene Scene
var SceneMap map[string]Scene = make(map[string]Scene)

func initScenes() {
	SceneMap["MenuMain"] = &MenuMain{
		Background: loadPng("img/test.png"),
	}
	CurrScene = SceneMap["MenuMain"]
}
