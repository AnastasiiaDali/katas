Given a string S, an integer L and an alignment A; fit S into an empty string of length L, aligning it to the left, right or centre according to A.
Fill the empty spaces with a dot '.'

e.g. if S is “Akash” and L=9

“Akash”, 9, L  gives "Akash...."  // aligned to the left
”Akash”, 9, R  gives "....Akash" // aligned to the right
”Akash”, 9, C  gives "..Akash.." // aligned in the centre


When centre alignment is not exactly possible, leave more spaces to the left.

e.g. ”Akash”, 10, C  gives "...Akash.."

L will always be >= length(S)