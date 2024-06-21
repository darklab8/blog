package trades

/*
Not implemented but found different materials

I found those scientist papers explaning it and showing different pseudocodes
- https://cse.buffalo.edu/faculty/miller/Courses/CSE633/Asmita-Gautam-Spring-2019.pdf
- https://moorejs.github.io/APSP-in-parallel/
- https://en.wikipedia.org/wiki/Parallel_all-pairs_shortest_path_algorithm

- I found this version in C. with code example/tests how to launch
	- https://github.com/domdicarlo/parallel-floyd-warshall/blob/master/src/floyd_warshall.c
- this version in Java. but no code examples how to work with it
	- https://rma350.github.io/2012/06/13/all-pairs-shortest-path-in-parallel-with-floyd-warshall-in-java.html
	- https://www.baeldung.com/java-executor-service-tutorial
	- https://algs4.cs.princeton.edu/44sp/FloydWarshall.java.html sequential floyd in java. should help to figure out how to work with parallel one in java.
- This version in Java, looks very dubious, as it is mot matching scientist papers
	- https://github.com/RJBrodsky/FloydWarshall-Sequential-vs-Parallel/blob/master/src/FloydWarshall.java
	- run with `javac FloydWarshall.java && java FloydWarshall`
- this version in Cuda language (first time seeing such language)
	- https://github.com/koallen/parallel-floyd-warshall

Searching places:
- seek github https://github.com/topics/floyd-warshall?o=asc&s=updated for floyd warshall and parallel floyd warshall
- google scientist papers
- try to find leetcode solutions

Conclusion onto Floyd Warshall parallelization.
Algorithms pretty much almost dont exist, and those that do, are very dubious if correct ones
and for some reason not matching scientist papers.
You need a good understanding of Floyd Warshall parallelization from scientist papers, in order to that in a right way.
*/
