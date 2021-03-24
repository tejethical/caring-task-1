<h1>Go Lang Interview Task</h1>

<h2>Description:</h2>
This program takes a txt file containing n number of strings separated by a new line as input and
sorts them using a distributed network of 4 nodes (master included) and prints the resulting
sorted array of strings in console.

The single binary file "caring-task-1" can be deployed as master or worker depending on the argument

<h2>Usage:</h2>
Run as master node
    
    caring-task-1 master <FILEPATH>
    
FILEPATH is the path of the input file

Run as worker node

    caring-task-1 worker <PORT_NUMBER>
    
PORT_NUMBER can be 91, 92 and 93



