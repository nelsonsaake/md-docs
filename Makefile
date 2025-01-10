.PHONY: ncommit run
.SILENT:
.ONESHELL:

ncommit:
	cls
	git add . 
	git commit -m "chore: commit everything"
	git push origin main
  
run:
	cls
	go run . 
  
reset:
	cls
	go run . run:api setup

env:
	cls
	env 

box:
	cls
	go run . gen:box
	
get:
	go get github.com/nelsonsaake/go@latest
	go get -u gorm.io/gorm

plg:
	cls
	cd ./src/kernel/plg
	go run .