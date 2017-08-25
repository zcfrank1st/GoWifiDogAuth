clean:
	rm -rf ../bin ../pkg

compile:
	go build -o ../bin/app bootstrap/bootstrap.go

all: clean compile
	mv ../bin/app /opt
	mv app.ini /opt
