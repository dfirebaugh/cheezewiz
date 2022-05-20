package vm

import (
	"github.com/robertkrimen/otto"
)

type scene interface {
	// SetPlayer(player gohan.Entity)
}

func Build(s scene) *otto.Otto {
	vm := otto.New()

	// vm.Set("addPlayerEntity", func(call otto.FunctionCall) otto.Value {
	// 	println(call.Argument(0).ToInteger())

	// player := entity.NewPlayer()
	// s.SetPlayer(player)

	// println(player)
	// entityID, _ := vm.ToValue(player)
	// return entityID
	// })

	// vm.Set("addPlayerControlSystem", func(call otto.FunctionCall) otto.Value {
	// 	gohan.AddSystem(system.NewPlayerControl())
	// 	return otto.Value{}
	// })
	// vm.Set("addMovementSystem", func(call otto.FunctionCall) otto.Value {
	// 	gohan.AddSystem(system.NewMovement())
	// 	return otto.Value{}
	// })
	// vm.Set("addRenderSystem", func(call otto.FunctionCall) otto.Value {
	// 	gohan.AddSystem(system.NewRenderer())
	// 	return otto.Value{}
	// })
	// vm.Set("addAnimationSystem", func(call otto.FunctionCall) otto.Value {
	// 	gohan.AddSystem(system.NewAnimator())
	// 	return otto.Value{}
	// })

	return vm
}

func Run(vm *otto.Otto, script string) {
	vm.Run(script)
}
