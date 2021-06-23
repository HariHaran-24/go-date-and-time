FROM golang:1.15.7-buster
RUN git clone https://github.com/HariHaran-24/go-date-and-time.git
WORKDIR go-date-and-time
CMD go run app.go
EXPOSE 8080
