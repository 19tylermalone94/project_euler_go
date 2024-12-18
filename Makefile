build:
	go build -o euler main.go $(wildcard p*.go)

run:
	./euler $(problem)

new-problem:
	@read -p "Enter problem number (e.g., 001): " num; \
	name="p$${num}.go"; \
	if [ -f "$$name" ]; then \
	  echo "Error: $$name already exists."; \
	else \
	  printf "package main\n\nimport \"fmt\"\n\nfunc init() {\n  problems[\"p$${num}\"] = problem$${num}\n}\n\nfunc problem$${num}() {\n  fmt.Println(\"Running Problem $$num\")\n  // solution here\n}\n" > "$$name"; \
	  echo "Created $$name"; \
	fi
