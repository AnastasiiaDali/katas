We're using to expressions looking like a+b where a and b are operands and + is the operation. This form is called the infix form because the operator is in between the operands.

A postfix notation of the same expression looks like ab+. More complex operations look like abcd+-*, which stands for a * {b - (c+d)}

For example,
82/ will evaluate to 4 (8 / 2)

138*+ will evaluate to 25 (1 + 8*3)

545*+5/ will evaluate to 5 ((5 + 4*5) / 5)

Using the stack data structure, find the value of any given postfix expression.

Assume that the postfix expression contains only single-digit numeric operands, without any whitespace, and the operations are +, -, * and /