# simple-go-app
A simple go app to use for bench marcking



--------

Simple go app

https://github.com/GrantZheng/kit/blob/master/docs/qiuck-start/creating_a_todo_app_using_gokit-cli.md following this as a tut

* Go mod init github.com/maxsimmonds1337/simple-go-app
* go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
* go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest 
* go install github.com/GrantZheng/kit@latest (installs https://github.com/GrantZheng/kit)
* kit new service vote (makes a new service) (maybe I should have done this outside of the simple go app, that would have helped with the two go.mod files but let’s push on)
    * I delete the generated go.mod (there seems to be issues with modules in modules, and it’s best practise not to do this)
    * Added a io.go file in a io folder under /vote/pkg to define a vote
* Then I defined the end points by adding them into the service.go file
* kit g s vote -w (this needs to be run in the root folder, --w generate some default service middleware, g is generate s is service and vote is the name of the service)
* Go mod tidy
* Go run vote/cmd/main.go (this now runs, sets up some stuff)
* I added a db folder under pkg and a do.go file that hold the db connection stuff
* I then added the business logic into the get/add functions (makeGetHandler)
* I added this to my basic
export DB_HOSTNAME=localhost
export DB_PORT=50500
export DB_DATABASE=hotdogs
export DB_USERNAME=db2inst1
export DB_PASSWORD=hotdog

export CGO_CFLAGS=-I/Users/max/go/pkg/mod/github.com/ibmdb/clidriver/include
export CGO_LDFLAGS=-L/Users/max/go/pkg/mod/github.com/ibmdb/clidriver/lib

* Remember to source the basic
* Then I install the ibmdb2 clidriver:
* cd ~/go/pkg/mod/github.com/ibmdb/go_ibm_db@v0.4.3/installer 
* Go run setup.go
* Cd ../../ (check that I have folders like include there) pwd to get the dir needed for the exports above (/Users/max/go/pkg/mod/github.com/ibmdb/clidriver)
* I have issues running this on m1/m2, there’s no ibmdb compiler for apple silicon yet, so I run:
* kit g d (kit generate docker) which generates me a docker file. It uses watcher so you don’t have to recompile the code each time you change the source
* 


Setup DB2 in Docker (see https://www.ibm.com/docs/en/db2/11.5?topic=SSEPGG_11.5.0/com.ibm.db2.luw.qb.server.doc/doc/t_install_db2CE_win_img.htm)

* Mkdir Docker
* Cd Docker
* Vim .env_list
* Add ~/Docker to you file sharing settings in docker (settings->resources)
￼
LICENSE=accept
DB2INSTANCE=db2inst1
DB2INST1_PASSWORD=hotdog
DBNAME=hotdogs
BLU=false
ENABLE_ORACLE_COMPATIBILITY=false
UPDATEAVAIL=NO
TO_CREATE_SAMPLEDB=false
REPODB=false
IS_OSXFS=false
PERSISTENT_HOME=false
HADR_ENABLED=false
ETCD_ENDPOINT=
ETCD_USERNAME=
ETCD_PASSWORD=
* docker pull ibmcom/db2 (or DOCKER_DEFAULT_PLATFORM=linux/amd64 docker pull ibmcom/db2 if using a Mac with m1/m2 chip)
* 
* docker run -h db2_hd --name db2_hd --restart=always --detach --privileged=true -p 50500:50500 --env-file .env_list -v ~/Docker:/database ibmcom/db2 (unfortunately, ibm doesn’t yet support m1/m2 so this will throw a warning)
* Docker logs <container name> Check the logs, it takes a while to setup for the first time 
* Enter into a terminal on the docker and run - su - db2inst1
* db2 connect to hotdogs
* db2 "CREATE TABLE votes (id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY, nominee VARCHAR(20), app VARCHAR(10))" (generate a table with auto inc id)
* Duo a test insert - db2 "INSERT INTO votes (nominee, app) VALUES ('Sandwich', 'Go')" (response should be something like this - DB20000I  The SQL command completed successfully.)
* [db2inst1@db2_hd ~]$ db2 "select * from votes"
* 
* ID          NOMINEE              APP       
* ----------- -------------------- ----------
*           1 Sandwich             Go        
* 
*   1 record(s) selected.


* Had to do this to get db2 connection working to dbvisulaiser
* db2 "update dbm cfg using svcename 50500"
* db2stop force
* db2start

docker network create my-BUNderful-network

docker build --platform linux/amd64  -t simple-go-app .
docker run --network my-Bunderful-network -p 8081:8081 -it simple-go-app

curl -X POST -H "Content-Type: application/json" -d '{
  "vote": {
    "id": 1,
    "nominee": "Taco",
    "app": "GO"
  }
}' http://localhost:8080/add


curl http://localhost:8081/get
