.DEFAULT_GOAL := build

build: 
	@echo "Compiling..."
	@make -s fumpt
	@go build -o ${CALC_EXE} ${CMD_FLDR}/${CALC_FLDR}/${MAIN_GO}
	@go build -o ${WFREQ_EXE} ${CMD_FLDR}/${WFREQ_FLDR}/${MAIN_GO}
	@go build -o ${CROSS_EXE} ${CMD_FLDR}/${CROSS_FLDR}/${MAIN_GO}
	@go build -o ${LOG_EXE} ${CMD_FLDR}/${LOG_FLDR}/${MAIN_GO}
	@echo "Executable files ready!\n"
	@echo "make calc - run Calculator"
	@echo "make wordfreq - run Most frequently occurring words"
	@echo "make slicecross - run Find common values in two crossing slices"
	@echo "make visitlog - run imitation of Visit log\n"
	@echo "tests and coverage - make test, make coverage"
	@echo "pkgsite documentation - make dvi"
	@echo "godoc dvi - make dvi_<exe_name>"
.PHONY: build

calc:
	@./${CALC_EXE}
.PHONY: calc

wordfreq:
	@./${WFREQ_EXE}
.PHONY: wordfreq

slicecross:
	@./${CROSS_EXE}
.PHONY: slicecross

visitlog:
	@./${LOG_EXE}
.PHONY: visitlog

test:
	go test -v ./${INT_FLDR}/...
.PHONY: test

clean:
	@rm -rf ${COVERAGE_REPORT}
	@rm -rf ${COVERAGE_REPORT}.html
	@rm -rf ${CALC_EXE}
	@rm -rf ${WFREQ_EXE}
	@rm -rf ${CROSS_EXE}
	@rm -rf ${LOG_EXE}