# Arithmetic-Calculator
This application implements a simple arithmetic expression parser that reads expressions from 'expression.txt' file and print the result. 
Parser uses a recursive descent approach to parse an expression and construct an abstract syntax tree (AST) from it. The expression can contain integers and operators such as +, -, *, / and parentheses.

Here are used grammar rules:
E -> T {+|- T}
T -> F {*|/ F}
F -> (E) | {0,1,2,3,4,5,6,7,8,9}

Finally, the Evaluate method is used to evaluate the value of the expression represented by the AST by recursively visiting the nodes in the tree and performing the operations specified by the nodes.

One can build and run application using Dockerfile.




