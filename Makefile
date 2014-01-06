link:
	@echo "Linking local copy to /usr/bin"
	@mv /usr/bin/mao /tmp/mao.bak
	@ln ./mao /usr/bin/mao
	@echo "Done"

unlink:
	@mv /tmp/mao.bak /usr/bin/mao
	@echo "Done!"
