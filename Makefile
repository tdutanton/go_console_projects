SHELL := /bin/sh

.DEFAULT_GOAL := build

INT_FLDR := internal
CMD_FLDR := cmd
CALC_FLDR := calculator
WFREQ_FLDR := wordfreq
CROSS_FLDR := slicecrossing
LOG_FLDR := visitlog

COVERAGE_REPORT := report
CALC_EXE := calc
WFREQ_EXE := wfreq
CROSS_EXE := slicecross
LOG_EXE := visitlog

MAIN_GO := main.go
CALC_PKG := calculator.go
WFREQ_PKG := wordfreq.go
CROSS_PKG := slicecrossing.go
LOG_PKG := visitlog.go

GOFUMPT := $(shell which gofumpt)
#PKGSITE := $(shell which pkgsite)

fmt:
	@go fmt ./...
.PHONY: fmt

lint:
	@golint ./...
.PHONY: lint

vet:
	@go vet ./...
.PHONY: vet

std_linters: fmt lint vet
.PHONY: std_linters

fumpt: ensure-fumpt std_linters
	@gofumpt -l -w .
.PHONY: fumpt

ensure-fumpt:
ifeq (, $(GOFUMPT))
	@echo "Gofumpt is not installed, installing..."
	@go install mvdan.cc/gofumpt@latest
	@echo "Gofumpt installed"
endif
.PHONY: ensure-fumpt

dvi_calc:
	@go doc -all ${INT_FLDR}/${CALC_FLDR}/${CALC_PKG}
.PHONY: dvi_calc

dvi_wordfreq:
	@go doc -all ${INT_FLDR}/${WFREQ_FLDR}/${WFREQ_PKG}
.PHONY: dvi_wordfreq

dvi_cross:
	@go doc -all ${INT_FLDR}/${CROSS_FLDR}/${CROSS_PKG}
.PHONY: dvi_cross

dvi_log:
	@go doc -all ${INT_FLDR}/${LOG_FLDR}/${LOG_PKG}
.PHONY: dvi_log

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
	@echo "dvi - make dvi_<exe_name>"
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

coverage:
	go test -v -cover -coverprofile=${COVERAGE_REPORT} ./${INT_FLDR}/...
	go tool cover -html=${COVERAGE_REPORT} -o ./${COVERAGE_REPORT}.html
.PHONY: coverage

clean:
	@rm -rf ${COVERAGE_REPORT}
	@rm -rf ${COVERAGE_REPORT}.html
	@rm -rf ${CALC_EXE}
	@rm -rf ${WFREQ_EXE}
	@rm -rf ${CROSS_EXE}
	@rm -rf ${LOG_EXE}

# ensure-pkgsite:
# ifeq (, $(PKGSITE))
# 	@echo "pkgsite is not installed, installing..."
# 	@go install golang.org/x/pkgsite/cmd/pkgsite@latest
# 	@echo "pkgsite installed"
# endif
# .PHONY: ensure-pkgsite

# dvi:
# 	pkgsite
# 	@echo "open http://localhost:8080 in your browser"
# .PHONY: dvi
