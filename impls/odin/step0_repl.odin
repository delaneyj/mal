package main

import "core:bufio"
import "core:fmt"
import "core:os"
import "core:strings"

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
	r: bufio.Reader
	buffer: [1024]u8
	bufio.reader_init(&r, os.stream_from_handle(os.stdin))
	defer bufio.reader_destroy(&r)

	for {
		defer free_all(context.temp_allocator)

		fmt.print("user> ")
		input, err := bufio.reader_read_string(&r, '\n', context.temp_allocator)
		if err != nil {
			fmt.println("Error reading input: ", err)
			continue
		}
		input = strings.trim_space(input)

		// if input == "exit\n" {
		// 	fmt.println("Exiting REPL.")
		// 	break
		// }

		fmt.println(rep(input))
	}
}
