STEPS = step0_repl
BINDGEN ?= bindgen

all: $(STEPS)

bindings:
	@$(BINDGEN) vendor/editline || { \
		echo ""; \
		echo "Could not run '$(BINDGEN)'."; \
		echo "Build/install odin-c-bindgen, or run:"; \
		echo "  make -C impls/odin BINDGEN=/path/to/bindgen.bin bindings"; \
		exit 127; \
	}

dist: mal

mal: step0_repl
	cp $< $@

$(STEPS): %: %.odin
	odin build $< -file -out:$@

clean:
	rm -f $(STEPS) mal

.PHONY: all clean bindings
