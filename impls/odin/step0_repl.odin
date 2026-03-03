package main

import "core:fmt"
import "core:mem"
import "vendor/editline"

read :: proc(input: string) -> string {
	return input
}

eval :: proc(tokens: string) -> string {
	return tokens
}

print :: proc(output: string) -> string {
	return output
}

rep :: proc(input: string) -> string {
	return print(eval(read(input)))
}

main :: proc() {
	repl_arena: mem.Dynamic_Arena
	mem.dynamic_arena_init(&repl_arena)
	defer mem.dynamic_arena_destroy(&repl_arena)

	repl_allocator := mem.dynamic_arena_allocator(&repl_arena)
	context.allocator = repl_allocator

	for {
		defer mem.dynamic_arena_reset(&repl_arena)

		line, ok := editline.readline("user> ", repl_allocator)
		if !ok {
			break
		}

		editline.add_history(line, repl_allocator)
		fmt.println(rep(line))
	}
}
