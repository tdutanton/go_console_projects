
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

dvi:
	@echo "To see the documentation open http://localhost:8080 in your browser"
	@pkgsite
.PHONY: dvi

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

coverage:
	go test -v -cover -coverprofile=${COVERAGE_REPORT} ./${INT_FLDR}/...
	go tool cover -html=${COVERAGE_REPORT} -o ./${COVERAGE_REPORT}.html
.PHONY: coverage

ensure-pkgsite:
ifeq (, $(PKGSITE))
	@echo "pkgsite is not installed, installing..."
	@go install golang.org/x/pkgsite/cmd/pkgsite@latest
	@echo "pkgsite installed"
endif
.PHONY: ensure-pkgsite

ensure-fumpt:
ifeq (, $(GOFUMPT))
	@echo "Gofumpt is not installed, installing..."
	@go install mvdan.cc/gofumpt@latest
	@echo "Gofumpt installed"
endif
.PHONY: ensure-fumpt
