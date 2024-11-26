package evaluator

import (
	"monkey/object"
)

var builtins map[string]*object.Builtin

func init() {
	initializeBuiltins()
}

func initializeBuiltins() {
	builtins = map[string]*object.Builtin{
		"len": &object.Builtin{
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1", len(args))
				}
				switch arg := args[0].(type) {
				case *object.Array:
					return &object.Integer{Value: int64(len(arg.Elements))}
				case *object.String:
					return &object.Integer{Value: int64(len(arg.Value))}
				default:
					return newError("argument to `len` not supported, got %s", args[0].Type())
				}
			},
		},
		"contains": &object.Builtin{
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 2 {
					return newError("wrong number of arguments. got=%d, want=2", len(args))
				}
				if args[0].Type() != object.ARRAY_OBJ {
					return newError("argument to `contains` must be ARRAY, got %s", args[0].Type())
				}
				arr := args[0].(*object.Array)
				target := args[1]

				for _, elem := range arr.Elements {
					if elem.Inspect() == target.Inspect() {
						return &object.Boolean{Value: true}
					}
				}
				return &object.Boolean{Value: false}
			},
		},
		"map": &object.Builtin{
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 2 {
					return newError("wrong number of arguments. got=%d, want=2", len(args))
				}
				if args[0].Type() != object.ARRAY_OBJ {
					return newError("argument to `map` must be ARRAY, got %s", args[0].Type())
				}
				if args[1].Type() != object.FUNCTION_OBJ {
					return newError("second argument to `map` must be FUNCTION, got %s", args[1].Type())
				}

				arr := args[0].(*object.Array)
				fn := args[1].(*object.Function)

				var newElements []object.Object
				for _, elem := range arr.Elements {
					result := applyFunction(fn, []object.Object{elem})
					if isError(result) {
						return result
					}
					newElements = append(newElements, result)
				}

				return &object.Array{Elements: newElements}
			},
		},
	}
}
