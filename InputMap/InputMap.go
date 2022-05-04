package InputMap

import (
	"github.com/faiface/pixel/pixelgl"
)

type InputMap struct {
	inputs []pixelgl.Button
}

var Inputs = InputMap{
	inputs: make([]pixelgl.Button, 0),
}

func (input *InputMap) Add(button pixelgl.Button) {
	input.inputs = append(input.inputs, button)
}

func (input *InputMap) Remove(button pixelgl.Button) {
	var index int
	found := false
	for i, in := range input.inputs {
		if in == button {
			index = i
			found = true
			break
		}
	}
	if found {
		input.inputs = append(input.inputs[:index], input.inputs[index+1:]...)
	}
}

func (input *InputMap) Clear() {
	input.inputs = input.inputs[:0]
}

func (input *InputMap) Contains(button pixelgl.Button) bool {
	for _, in := range input.inputs {
		if in == button {
			return true
		}
	}
	return false
}
