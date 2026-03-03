package editline

when ODIN_OS == .Linux {
	@(export) foreign import lib "system:libedit.so.2"
} else when ODIN_OS == .Darwin {
	@(export) foreign import lib "system:edit"
} else {
	@(export) foreign import lib "system:edit"
}
