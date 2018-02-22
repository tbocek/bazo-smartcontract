TARGETBAEPRINT=Archbot-Technischer_Bericht_eprint_ba.pdf
TARGETBAPRINT=Archbot-Technischer_Bericht_print_ba.pdf
LL=latexmk -halt-on-error -pdf -pdflatex="xelatex" -bibtex
CLEAN=latexmk -C

BAE: $(TARGETBAEPRINT)
BA: $(TARGETBAPRINT)

all: BAE
.PHONY : $(TARGETBAEPRINT) $(TARGETBAPRINT) clean all

$(TARGETBAEPRINT): $(TARGETBAEPRINT:%.pdf=%.tex)
	$(LL) $<
$(TARGETBAPRINT): $(TARGETBAPRINT:%.pdf=%.tex)
	$(LL) $<

clean:
	$(CLEAN)