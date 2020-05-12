<h1>Go TODO</h1>
Start GRPC server: ./server -grpc-port=9090 -db-host=localhost:3306 -db-user=root -db-password=root -db-schema=todo -http-port=8080</br>
Call grpc api: ./client-grpc -server=localhost:9090</br>
 ./client-http -server=http://localhost:8080</br>