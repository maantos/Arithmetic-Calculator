# Arithmetic-Calculator
Simple arithmetic LL1 recuresive descent parserProgram that can parse and evaluate arithmetic expressions.


USed grammar:
E -> T {+|- T}
T -> F {*|/ F}
F -> (E) | {0,1,2,3,4,5,6,7,8,9}
