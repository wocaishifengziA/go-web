msg = $(m)
push:
	git add .
	git commit -m $(msg)
	git push

t:
	go run base/main.go