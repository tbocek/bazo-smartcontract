TARGET=Template.pdf
LL=latexmk -halt-on-error -pdf -pdflatex="xelatex" -bibtex
CLEAN=latexmk -c

all: $(TARGET)
.PHONY : $(TARGET) clean all

$(TARGET): $(TARGET:%.pdf=%.tex)
	$(LL) $<

clean:
	$(CLEAN)
