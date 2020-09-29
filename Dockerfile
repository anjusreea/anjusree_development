FROM ubuntu
#MAINTAINER anjusree.anju88@gmail.com

RUN apt-get update

RUN apt-get install -y go

COPY employee.go /root/employee.go

CMD ["go","/root/employee.go"]
