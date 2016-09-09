package scene

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	systems "github.com/engoengine/temod/Systems"
)

type MyScene struct{}

// Type uniquely defines your game type37G
func (*MyScene) Type() string { return "myGame" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*MyScene) Preload() {
	engo.Files.Load("textures/kiwi.svg")
	engo.Files.Load("textures/Mushroom2.png")
	engo.Files.Load("textures/bayless.jpg")
	engo.Files.Load("textures/city.png")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*MyScene) Setup(world *ecs.World) {
	kbspeed := float32(400)
	engo.Input.RegisterButton("AddCity", engo.F1)
	common.SetBackground(color.White)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})
	world.AddSystem(&systems.CityBuildingSystem{})
	world.AddSystem(common.NewKeyboardScroller(kbspeed, engo.DefaultHorizontalAxis, engo.DefaultVerticalAxis))
	world.AddSystem(&common.EdgeScroller{400, 20})
	world.AddSystem(&common.MouseZoomer{-0.125})

}
