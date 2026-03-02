STEPS = step0_repl

all: $(STEPS)

dist: mal

mal: step0_repl
	cp $< $@

$(STEPS): %: %.odin
	odin build $< -file -out:$@

clean:
	rm -f $(STEPS) mal
