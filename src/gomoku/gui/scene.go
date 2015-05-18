package gui

type Scene interface {
	PlayScene()
}

var CurrScene *Scene
var SceneMap map[string]*Scene = make(map[string]*Scene)
