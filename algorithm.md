Case I

A adds 3 goods a1, a2, a3
B adds 3 goods b1, b2, b3 and gets offer a1, a2, a3
C adds 3 goods c1, c2, c3 and get offer a1, a2, a3, b1, b2, b3

A when opens next time gets Offers b1, b2, b3, c1, c2, c3 and he likes b2, c1
B when opens next time gets offers a1r, a2r, a3r, c1, c2, c3 and he likes a2r (subscript r means reffered). cycle is formed B now has to deliver b2 to A
A get notified and has to deliver a2 to B

Case II

A adds 3 goods a1, a2, a3
B adds 3 goods b1, b2, b3 and gets offer a1, a2, a3
C adds 3 goods c1, c2, c3 and get offer a1, a2, a3, b1, b2, b3

B when opens next time gets offers a1, a2, a3, c1, c2, c3 and he likes c2
A when opens next time gets Offers b1, b2, b3, c1, c2, c3 and he likes b2
B when opens next time gets offers a1r, a2r, a3r, c1, c2, c3 and as he doesnt like any good of a he passes Ar to C
C when opens next time gets Offers b1r, b2r, b3r, a1r, a2r, a3r and he likes a1r and cycle is formed. C has to deliver c2 to B
B get notified and he has deliver b2 to A and propagate chain event to A.
A get notified and he has deliver a1 to C.
