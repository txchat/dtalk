build_amd:
	@echo '┌ start	build amd64'
	@bash ./script/build/multi-service.sh
	@echo '└ end	build amd64'

build_arm:
	@echo '┌ start	build 	arm64'
	@bash ./script/build/multi-service.sh arm64
	@echo '└ end	build 	arm64'

# 编译二进制 amd64
quick_build_amd:
	@echo '┌ start	quick build amd64'
	@bash ./script/build/quick_build.sh
	@echo '└ end	quick build amd64'

# 编译二进制 amr64
quick_build_arm:
	@echo '┌ start	quick build arm64'
	@bash ./script/build/quick_build.sh arm64
	@echo '└ end	quick build arm64'