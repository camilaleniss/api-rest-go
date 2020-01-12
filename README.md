# api-rest-go

Tecnologies: <br>
DataBase: CockroachDB<br>
FrontEnd: Vue.js and bootstrap-vuejs<br>
Backend: Go<br>
<br>

Endpoints Backend:<br>
GET the domains searched<br>
localhost:8082/api/domains<br>
GET the info of an specific domain<br>
localhost:8082/api/{id}<br>

<br>
Database run in port localhost:8080<br>

<br>
Frontend runs in localhost:8081<br>

<h2>Step by step SetUp</h2><br>
To run the node of the database run in one terminal the following command in the folder where is cockroach.exe<br>
./cockroach.exe start --insecure --listen-addr=localhost <br><br>
Then open other terminal and do the following commands<br>
./cockroach.exe sql --insecure --host=localhost:26257<br>
CREATE USER IF NOT EXISTS maxroach; <br>
CREATE DATABASE domains;<br>
CREATE TABLE IF NOT EXISTS domain (host STRING PRIMARY KEY, ssl_grade STRING, ssl_previous_grade STRING, last_search TIMESTAMPTZ);<br>
<br><br>
To run the frontend open other terminal in api-rest-go/gui/domainapp and execute:<br>
npm run serve<br>


