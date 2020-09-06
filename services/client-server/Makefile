docker-%:
	@docker build -t alextanhongpin/$* -f Dockerfile.$*  .

server:
	@docker run -d -p 8080:8080 alextanhongpin/server
