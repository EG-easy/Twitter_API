# twitter_textmining
Twitter word scaning tool to find trend words

# Installation
- mongoDB

You can check official webpage!

https://www.mongodb.com/download-center?jmp=nav#community

# Quick Start!
0. git clone this repository
```shell
$ git clone https://github.com/EG-easy/twitter_textmining.git
```

1. get Twitter API from https://apps.twitter.com and paste them to setup.sh file

2. reflect the API setting
```shell
$ source setup.sh
```

3. start the setting of mongoDB 
```shell
# open another shell
$ nsqlookupd

# open another shell
$ nsqd --lookupd-tcp-address=localhost:4160

# open another shell
$ sudo mongod --dbpath /var/lib/mongodb --logpath /var/log/mongodb.log  # set the database/log files first

# open another shell
$mongo
```

4. build counter file
```shell
$ cd counter
$ go build -o counter
$ ./counter
```

5. build twitter API file
```shell
$ cd twittervotes
$ go build -o twittervotes
$ ./twittervotes
```

6. build api file
```shell
$ cd api
$ go build -o api
$ ./api
```

7. build web file
```shell
$ cd web
$ go build -o api
$ ./web
```

8. access to http://localhost:8081

Now you can search any words used in twitter and compare trending words!
