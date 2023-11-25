push:
	git add .
	git commit -m "ЭРИК ПИДОРАС"
	git push -u origin golang

merge:
	git checkout main
	git merge golang
	git push -u origin main