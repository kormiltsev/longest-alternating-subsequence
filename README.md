# Longest alternating subsequence

Incomig sequence is like a(1),a(2),a(3),...a(n) 
outcome a(l1),a(l2),a(l3),..a(i)... where l1<l2<l3 and for any i a(i-1)<a(i)>a(i+1) or a(i-1)>a(i)<a(i+1)

Find subsequense as long as possible

(i) as small as possible

0<a(k)<1000000000
0<n<1000

# Target:
time<0.5 sec

memory<64 MB

income: standart input or input.txt

Like: "n,a1,a2,..." 

outcome: standart outcome or output.txt

ex:

input 
123343525436

output 
14352546

# solution 1 (from file)
for use input.txt:
go run ./cmd/task131022/main.go

to generate new input.txt file there is generatefile.go:
go run ./files/generatefile.go

can be modified to work with stdin

use ./files/generatefile.go to create new random ./files/input.txt

# solution 2 (from stdin)
Fast and no memory. Calculation on the fly, stdout as ready. Using channels.

Upload data from stdin. There is generator you can use:
go run ./files/generatorstdin.go | go run ./cmd/task131022/mainstdin.go 