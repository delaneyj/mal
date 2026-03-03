package editline

import "core:c"
import "core:c/libc"
import "core:strings"

@(default_calling_convention="c")
foreign lib {
	@(link_name="readline")
	_readline :: proc(prompt: cstring) -> cstring ---
	@(link_name="add_history")
	_add_history :: proc(line: cstring) -> c.int ---
	@(link_name="read_history")
	_read_history :: proc(filename: cstring) -> c.int ---
	@(link_name="write_history")
	_write_history :: proc(filename: cstring) -> c.int ---
}

readline :: proc(prompt: cstring = "user> ", allocator := context.allocator) -> (line: string, ok: bool) {
	c_line := _readline(prompt)
	if c_line == nil {
		return "", false
	}

	line = strings.clone_from_cstring(c_line, allocator)
	libc.free(rawptr(c_line))
	return line, true
}

add_history :: proc(line: string, allocator := context.temp_allocator) {
	c_line := strings.clone_to_cstring(line, allocator)
	_add_history(c_line)
}

read_history :: proc(filename: string, allocator := context.temp_allocator) {
	c_filename := strings.clone_to_cstring(filename, allocator)
	_read_history(c_filename)
}

write_history :: proc(filename: string, allocator := context.temp_allocator) {
	c_filename := strings.clone_to_cstring(filename, allocator)
	_write_history(c_filename)
}
