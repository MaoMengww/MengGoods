.PHONY: start-all stop-all tidy

# 一键启动
start-all:
	@echo "Starting the engine..."
	@chmod +x script/start_all.sh
	@./script/start_all.sh

