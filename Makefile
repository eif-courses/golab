include .env

stop_containers:
	@echo "Stopping containers!"
	for /f %%i in ('docker ps -q') do ( \
		echo "Stopping container %%i"; \
		if docker stop %%i 2>NUL; then \
			echo "Container %%i stopped successfully"; \
		else \
			echo "Container %%i not found or already stopped!"; \
		fi \
	)