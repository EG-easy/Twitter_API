# twitter_textmining
Twitter word scaning tool to find trend words

# installation
- mongoDB

You can check official webpage!
https://www.mongodb.com/download-center?jmp=nav#community

# Quick Start!
0. git clone this repository
```shell:~
$ git clone https://github.com/EG-easy/twitter_textmining.git
```
1. get Twitter API from apps.twitter.com

2. reflect the API setting
```shell:~/twitter_textmining/
$ source setup.sh
```

3. start the setting of mongoDB 
```shell
# open another shell
$ nsqlookupd
$ nsqd --lookupd-tcp-address=localhost:4160
# open another shell
$ sudo mongod --dbpath /var/lib/mongodb --logpath /var/log/mongodb.log
# open another shell
$mongo
```

4. 
```shell-session:~

```

5. 
```shell-session:~
$ mongod
```


