Comparison, profiling and optimizing with paralleling.

## Intro

During the development of [a game data viewing tool for a space simulator]({{.DarkstatUrl}}), I encountered a need to calculate all shortest paths between 2218 vertices, having 30250 edges in a directional graph. That is required to calculate profit per time, for different trading routes in a game of flying between star systems and trading. ^_^.

## Floyd All Shortest Paths

I started with a simple algorithm of a Floyd Warshal to calculate all paths. By taking [page from wikipedia about it](https://en.m.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm) I implemented it in golang and measured the results

floyd.go
```go
{{.FloydMain}}
```

floyd_test.go: Unit test for launch
```go
{{.FloydTest}}
```

It showed too slow calculations on the desired graph size mentioned before regretfully
```
floyd:
floyder.GetDist("li01_01_base", "li01_to_li02")= 22154.047859719714
floyder.GetDist("li01_to_li02", "li02_to_li01")= 0
floyder.GetDist("li02_to_li01", "li12_02_base")= 9239.98628904769
GetDist(graph, dist, "li01_01_base", "li01_02_base") 21397.06209135834
GetDist(graph, dist, "li01_01_base", "br01_01_base") 128222.95301963031
GetDist(graph, dist, "li01_01_base", "li12_02_base") 31394.034148767405
time=2024-06-16T12:29:47.108+02:00 level=DEBUG msg="time_measure 2m20.534426392s | trade routes calculated"
```
Initially, it took more than 2 minutes and 20 seconds to calculate the necessary data with Floyd, which happened because of float64 usage. The time improved considerably to just 60 seconds if using integers instead of floats for matrix though. The example above is already of a fixed to integer usage algorithm.

The speed was not very satisfying, therefore it made me look further. Can I parallel this algorithm to speed up? [The found article](https://cse.buffalo.edu/faculty/miller/Courses/CSE633/Asmita-Gautam-Spring-2019.pdf) from a university hints at yes. However the algorithm looked dependent on previous parallelized jobs, it looked too complex, so I did not risk implementing it. Trustworthy code examples are rather tough to find for parallel Floyd too.

## Johnson's Algorithms

Looking through my options on [shortest paths algorithm wikipedia page](https://en.wikipedia.org/wiki/Shortest_path_problem), i saw that Johnson’s algorithm is very promising, as it is told to be possibly faster on sparse graph

I checked [geeks web site](https://www.geeksforgeeks.org/implementation-of-johnsons-algorithm-for-all-pairs-shortest-paths/) for inspiration how to implement, as it had already implementation in C++, Java, Python and Javascript. By rewriting the Java algorithm to Golang, I received a speed almost equal to Floyd's calculations with integers. (Initially, I thought it was twice faster since I discovered Integer optimization for Floyd later)

```
GetDist(graph, dist, "li01_01_base", "li01_to_li02")= 22148
GetDist(graph, dist, "li01_to_li02", "li02_to_li01")= 0
GetDist(graph, dist, "li02_to_li01", "li12_02_base")= 9235
GetDist(graph, dist, "li01_01_base", "li01_02_base") 21389
GetDist(graph, dist, "li01_01_base", "br01_01_base") 128172
GetDist(graph, dist, "li01_01_base", "li12_02_base") 31383
time=2024-06-16T12:33:57.060+02:00 level=DEBUG msg="time_measure 1m12.256099084s | trade routes calculated"
```

It was still a rather horrible time to calculate it. What i did do next? I turned on the default Golang profiling. By just adding extra code lines from std lib into the unit test beginning. And running after that `go tool pprof johnson.prof` and then inserting command `web`. That opened the browser page with the visual profile shown below.
```go
func TestTradeRoutesJohnson(t *testing.T) {
    //The code i don't wish to profile
	configs := configs_mapped.TestFixtureConfigs()
	graph := MapConfigsToFloyder(configs, WithFighterPaths(false))

	// for profiling only stuff.
	f, err := os.Create("johnson.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

    // Code i wish to profile
	timeit.NewTimerF(func(m *timeit.Timer) {
		johnson := NewJohnsonFromGraph(graph)
		var dist [][]int = johnson.Johnsons()

		fmt.Println(`GetDist(graph, dist, "li01_01_base", "li01_to_li02")=`, GetDist(graph, dist, "li01_01_base", "li01_to_li02"))
		fmt.Println(`GetDist(graph, dist, "li01_to_li02", "li02_to_li01")=`, GetDist(graph, dist, "li01_to_li02", "li02_to_li01"))
		fmt.Println(`GetDist(graph, dist, "li02_to_li01", "li12_02_base")=`, GetDist(graph, dist, "li02_to_li01", "li12_02_base"))
		dist1 := GetDist(graph, dist, "li01_01_base", "li01_02_base")
		dist2 := GetDist(graph, dist, "li01_01_base", "br01_01_base")
		dist3 := GetDist(graph, dist, "li01_01_base", "li12_02_base")
		fmt.Println(`GetDist(graph, dist, "li01_01_base", "li01_02_base")`, dist1)
		fmt.Println(`GetDist(graph, dist, "li01_01_base", "br01_01_base")`, dist2)
		fmt.Println(`GetDist(graph, dist, "li01_01_base", "li12_02_base")`, dist3)
		assert.Greater(t, dist1, 0)
		assert.Greater(t, dist2, 0)
		assert.Greater(t, dist3, 0)
	}, timeit.WithMsg("trade routes calculated"))
}
```

![]({{.StaticRoot}}shortest_paths/johnson_java_first.png)

I discovered that the provided Java algorithm had a problem with its `findMinDistanceVertex` function. After comparing the Java algorithm with C++ and Python algorithms ([at the same page]((https://www.geeksforgeeks.org/implementation-of-johnsons-algorithm-for-all-pairs-shortest-paths/))) I saw they use Priority Queue data structure, while Java just used search in a list. Suspecting not efficiency from it, I rewrote Dijkstra's algorithm to usage of Priority Queue, by taking info for it from [this golang std library](https://pkg.go.dev/container/heap) for heap ( Later was discovered that wiki page for Dijkstra ( https://en.wikipedia.org/wiki/Dijkstra's_algorithm ) actually documents both those versions and tells their complexity, the heap version is obviously better )

```
GetDist(graph, dist, "li01_01_base", "li01_to_li02")= 22148
GetDist(graph, dist, "li01_to_li02", "li02_to_li01")= 0
GetDist(graph, dist, "li02_to_li01", "li12_02_base")= 9235
GetDist(graph, dist, "li01_01_base", "li01_02_base") 21389
GetDist(graph, dist, "li01_01_base", "br01_01_base") 128172
GetDist(graph, dist, "li01_01_base", "li12_02_base") 31383
time=2024-06-16T14:02:39.834+02:00 level=DEBUG msg="time_measure 28.153478383s | trade routes calculated"
```
The results were promising. Now the algorithm calculated already in 28 seconds! But I wished faster, as it was still not enough for me. I wondered how to improve it?

Running the profiler again I saw all the slowness coming from calculating Dijkstra's algorithms as part of Johnson's algorithm.
![]({{.StaticRoot}}shortest_paths/johnson_sequential.png)

Checking the code I noticed rather an easy spot for parallelization, as calculations looked independent and used the same read-only data.
```go
func (g *Johnson) johnsons() [][]int {
        // ...Other Johnson Code
    var distances [][]int = make([][]int, g.vertices)

    for s := 0; s < g.vertices; s++ {
        distances[s] = g.dijkstra(s)
    }
        // ...Other Johnson Code
}
```

Could be the answer so simple? I checked it by adding simple Golang parallelism and received exactly the expected results, but in less than 6 seconds!  yay!
```
GetDist(graph, dist, "li01_01_base", "li01_to_li02")= 22148
GetDist(graph, dist, "li01_to_li02", "li02_to_li01")= 0
GetDist(graph, dist, "li02_to_li01", "li12_02_base")= 9235
GetDist(graph, dist, "li01_01_base", "li01_02_base") 21389
GetDist(graph, dist, "li01_01_base", "br01_01_base") 128172
GetDist(graph, dist, "li01_01_base", "li12_02_base") 31383
time=2024-06-16T14:27:25.075+02:00 level=DEBUG msg="time_measure 5.880306556s | trade routes calculated"
```

Providing the code for it:

heap.go
```go
{{.HeapCode}}
```

johnson.go
```go
{{.JohnsonCode}}
```

johnson_test.go
```go
{{.JohnsonTest}}
```

That was a success! With receiving still same values, the calculating was way faster just with the touch of a very simple Golang paralelism using go routines and channels to collect the result. We went from 2 minutes and 20 seconds (technically 1 minute) to the achieved 6 seconds results for 2218 vertixes, and 30250 edges, which was very satisfying and 34 times faster. [The same code can be found in Github folders](https://github.com/darklab8/blog/tree/master/blog/articles/article_detailed/article_shortest_paths/trades)

## Further optimizations

Found [the wiki page containing three different parallel methods for all shortest paths problem solving](https://en.wikipedia.org/wiki/Parallel_all-pairs_shortest_path_algorithm). The parallelizing Dijkstra solution made for Johnson's algo looks like the DijkstraAPSP described in it. There is room to try the advanced choices for Dijkstra parallelization (that will have potentially no gain), and Floyd parallelization.

Resources to find exact realizations for those different options are very scarce though. We could optimize by skipping calculations for all shortest paths originating from vertexes I use only as intermediate travel points, so we can replace for them Dijkstra calculations with a filled array `[Inf, Inf, Inf, 0(for source index), Inf, Inf, Inf...]` instead of calculating Dijkstra itself. The given task requires me to get distances starting from intermediate traveling points too though, *so I could not apply this optimization* for my requirements of the task.

Alternatively, there is another possible optimization. the given domain of data has as dense connections only specific star systems, with every vertex connected to each other, and connected very scarcely only by a few edges between star systems. We could have calculated rapidly all shortest paths for each system (even with floyd), and then we could have built a second graph that contains only "space bases" and jump gates/holes connecting star systems. That would have decreased amount of vertex from 2218 to less than 900 vertex and made faster potential total calculations. *But this optimization was not made too*, because there is a simple alternative.

The given domain of data inside each star system contains spaceship flight speed boosters called Trade Lanes. They are made out of multiple rings connected in a line. Originally each ring had its own vertex connected with edges to its neighbors. I applied optimization by joining intermediate vertexes in a single "trade lane", and thus the amount of vertex decreased from 2218 to 1218 without almost no drop in the quality of calculations.

Additionally, optimization was made by removing Johnson's algorithm parts. Since the graph actually had no edges with negative weights... Johnson's Algorithm was not really required. So it was stripped down to DijkstraAPSP already written inside of it.

This final optimization gave me `1.5 seconds total time` for calculation runs to get all the necessary shortest trading route distances between space bases in a galaxy of a space simulator, and that is a satisfying end result for 1218 vertices, having 22896 edges in a directed graph.

Since we used DijkstraAPSP, it was possible to modify the algorithm for returning reconstructed shortest exact paths. That was actually very desirable for the task goals, and thankfully possible. Modifications were made based on [wiki article for Dijkstra algo](<https://en.wikipedia.org/wiki/Dijkstra%27s_algor>)

Providing the most final DijkstraAPSP algorithm, calculating all shortest paths through DijkstraSSSP in Golang, with the provided ability for path reconstructions.

graph.go
```go
{{.GraphCode}}
```

dijkstra_apsp.go
```go
{{.DijkstraApsp}}
```

dijkstra_apsp_test.go
```go
{{.DijkstraApspTest}}
```

TestDijkstraAPSPWithGraph output:
```
solved matrix
      0      5      9      8
    INF      0      4      3
    INF    INF      0    INF
    INF    INF      1      0
get dists and paths
a -> c =  8 path= [{ a -1 0 0} {a b 0 1 5} {b c 1 3 3} {c  3 -1 0}]
a -> b =  5 path= [{-1 0 0} {0 1 5} {1 -1 0}]
a -> d =  9 path= [{-1 0 0} {0 1 5} {1 3 3} {3 2 1} {2 -1 0}]
detailed path reconstruction, useful for debug
prev=                    a next=                    b prev_node=     0 next_node=     1 dist=     5
prev=                    b next=                    c prev_node=     1 next_node=     3 dist=     3
PASS
```

Attachments:
- [the galaxy map for the same data for which we built the graph](https://fifthbarrier.github.io/Navmap/#q=)
- [Code examples from this article, cleaned from domain-specific stuff](https://github.com/darklab8/blog/tree/master/blog/articles/article_detailed/article_shortest_paths/trades)
- [Code in use for space simulator trading routes calculations](https://github.com/darklab8/fl-configs/tree/master/configs/configs_export/trades)
- [This article page in repository of static site generator](https://github.com/darklab8/blog/blob/master/blog/articles/article_detailed/article_shortest_paths/article.md).

All examples have unit tests on how to use them ^_^. The links above may contain a version of the article and code examples more up-to-date or with other fixes. The research results were applied to [darkstat](<https://github.com/darklab8/fl-darkstat>) tool for trading routes calculations between space bases in the [Freelancer Discovery](https://discoverygc.com/) modding community.

DarkStat trading routes tab:
![]({{.StaticRoot}}shortest_paths/darkstat_tab_trades.png)
