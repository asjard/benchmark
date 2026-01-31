servers ?= bench_asjard \
			bench_beego \
			bench_default \
		 	bench_echo \
			bench_fasthttp \
			bench_gin \
			bench_go_chassis \
			bench_go_zero \
			bench_kratos


.PHONY: update
update: $(servers) ## 更新服务


.PHONY: $(servers)
$(servers):
	servers=$@ make do_$(MAKECMDGOALS)

.PHONY: do_update
do_update:
	cd $(PWD)/servers/$(servers) && go get -u
