docker:
	docker build -t apartmentservice .
	docker run -p 9090:9090 apartmentservice
