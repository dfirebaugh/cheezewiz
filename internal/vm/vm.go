package vm

import (
	"cheezewiz/internal/ecs/entity"
	"cheezewiz/internal/ecs/system"

	"code.rocketnine.space/tslocum/gohan"
	"github.com/robertkrimen/otto"
)

type scene interface {
	SetPlayer(player gohan.Entity)
}

func Build(s scene) *otto.Otto {
	vm := otto.New()

	vm.Set("addPlayerEntity", func(call otto.FunctionCall) otto.Value {
		println(call.Argument(0).ToInteger())

		s.SetPlayer(entity.NewPlayer())
		return otto.Value{}
	})

	vm.Set("addPlayerControlSystem", func(call otto.FunctionCall) otto.Value {
		gohan.AddSystem(system.NewPlayerControl())
		return otto.Value{}
	})
	vm.Set("addMovementSystem", func(call otto.FunctionCall) otto.Value {
		gohan.AddSystem(system.NewMovement())
		return otto.Value{}
	})
	vm.Set("addRenderSystem", func(call otto.FunctionCall) otto.Value {
		gohan.AddSystem(system.NewRenderer())
		return otto.Value{}
	})
	vm.Set("addAnimationSystem", func(call otto.FunctionCall) otto.Value {
		gohan.AddSystem(system.NewAnimator())
		return otto.Value{}
	})

	return vm
}

func Run(vm *otto.Otto, script string) {
	vm.Run(script)
}
