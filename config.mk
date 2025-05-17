SHELL := /bin/sh

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
PKGSITE := $(shell which pkgsite)