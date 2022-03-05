HARNESSES := go rust
TARGETS := deps.install deps.update deps.upgrade verify verify.fix generate test test.unit test.smoke test.report test.unit.report test.smoke.report

define make-target
## Run $1 test harness target $2
testharness.$1.$2:
	cd ./test-harnesses/$1 && ${MAKE} $2
.PHONY: testharness.$1.$2
endef

$(foreach harness,$(HARNESSES),$(foreach target,$(TARGETS),$(eval $(call make-target,$(harness),$(target)))))
