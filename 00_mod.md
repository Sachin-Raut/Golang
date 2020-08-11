How to create/add modules
(dependency management)


1. create a repo on github (for eg. "bookstore-users-api") 
2. copy its https path. Click on "Code" button, and there will be "SSH" & "HTTPS" links. 
For "bookstore-users-api" github repo, my https path is as follows
 (https://github.com/Sachin-Raut/bookstore-users-api.git)
3. Lets say u want to create a "bookstore-users-api" on Desktop
4. In terminal(if u r using Mac) or command prompt(if u r using Windows), navigate to Desktop. Clone the repo using following command.
5. $ git clone https://github.com/Sachin-Raut/bookstore-users-api.git
6. In terminal(if u r using Mac) or command prompt(if u r using Windows), navigate to "bookstore-users-api".
7. Now execute the following command
8. $ go mod init github.com/Sachin-Raut/bookstore-users-api 
9. This automatically generates "go.mod" file in your local project file (Desktop/bookstore-users-api )
10. Now if you open your project, your IDE will show you "go.mod"
11. Now when you any new dependency in your project, "go.mod" file will keep track of that dependency.
12. And 1 more file "go.sum" will be created when you add new dependency to the project.
13. Remember never touch/edit "go.sum" file.
14. For beginner this info is enough to start a new Go project.