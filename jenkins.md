## this is a jenkins tutorial for java project

1. install jenkins

* 1.1 [Download Jenkins](http://mirrors.jenkins.io/war-stable/latest/jenkins.war)
* 1.2 Open up a terminal in the download directory and run java -jar jenkins.war
* 1.3 Browse to http://localhost:8080 and follow the instructions to complete the installation. 

2. install (Maven Integration plugin)

3. crate a maven job

4. configure for this job

4.1 general

    set project name and description
   
4.2  source control

    git
 
    set repository url and credentials are jenkins account
 
    Branches to build    master

4.3  trigger

    set Poll SCM
    schedule   H/2 * * * *  
    
4.4  build

    pom.xml

4.5  Post Steps

    install publish over SSH  plugin
    configure ssh

    set ssh server and transfer

transfer:

    set source files:  target/*.jar
    remove prefix:  target/
    remote directory:  /Users/renhuo/jenkinsDeploy
    Exec command:
        #!/bin/sh
        echo 'start to java project'
        cd  /Users/renhuo/jenkinsDeploy
        java -jar javaTest2-1.0-SNAPSHOT.jar 
        echo 'complete java project'
   
 
[doc:](https://jenkins.io/doc/)
[user doc:](http://m.blog.csdn.net/Evankaka/article/details/50518959)
   