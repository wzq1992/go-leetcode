本周作业
----------
1. [转换成小写字母（Easy）](https://leetcode-cn.com/problems/to-lower-case/)

    * [作业代码](homework/lc709_to_lower_case.go)

2. [最后一个单词的长度（Easy）](https://leetcode-cn.com/problems/length-of-last-word/)

    * [作业代码](homework/lc58_length_of_last_word.go)

3. [宝石与石头（Easy）](https://leetcode-cn.com/problems/jewels-and-stones/)

    * [作业代码](homework/lc771_jewels_and_stones.go)

4. [字符串中的第一个唯一字符（Easy）](https://leetcode-cn.com/problems/first-unique-character-in-a-string/)

    * [作业代码](homework/lc387_first_unique_character_in_a_string.go)

5. [最长公共前缀（Easy）](https://leetcode-cn.com/problems/longest-common-prefix/)

    * [作业代码](homework/lc14_longest_common_prefix.go)

6. [反转字符串（Easy）](https://leetcode-cn.com/problems/reverse-string/)

    * [作业代码](homework/lc344_reverse_string.go)

7. [翻转字符串里的单词（Medium）](https://leetcode-cn.com/problems/reverse-words-in-a-string/)
    
    * [作业代码](homework/lc151_reverse_words_in_a_string.go)

8. [仅仅反转字母（Easy）](https://leetcode-cn.com/problems/reverse-only-letters/)

    * [作业代码](homework/lc917_reverse_only_letters.go)


课上例题
----------
1. [网络延迟时间（Medium）](https://leetcode-cn.com/problems/network-delay-time/)

    * [练习代码](./lessons/lc743_network_delay_time.go)

2. [阈值距离内邻居最少的城市（Medium）](https://leetcode-cn.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/)

    * [练习代码](./lessons/lc1334_find_the_city_with_the_smallest_number_of_neighbors_at_a_threshold_distance.go)

3. [连接所有点的最小费用（Medium）](https://leetcode-cn.com/problems/min-cost-to-connect-all-points/)

    * [练习代码](./lessons/lc1584_min_cost_to_connect_all_points.go)

4. [字符串转换整数 (atoi) （Medium）](https://leetcode-cn.com/problems/string-to-integer-atoi/)

    * [练习代码](./lessons/lc8_string_to_integer_atoi.go)

5. [实现 strStr() （Easy）](https://leetcode-cn.com/problems/implement-strstr/)

    * [练习代码](./lessons/lc28_implement_strstr.go)

6. [验证回文串（Easy）](https://leetcode-cn.com/problems/valid-palindrome/)

    * [练习代码](./lessons/lc125_valid_palindrome.go)

7. [验证回文字符串 Ⅱ（Easy）](https://leetcode-cn.com/problems/valid-palindrome-ii/)

    * [练习代码](./lessons/lc680_valid_palindrome_ii.go)

8. [最长回文子串（Medium）](https://leetcode-cn.com/problems/longest-palindromic-substring/)

    * [练习代码](./lessons/lc5_longest_palindromic_substring_naive.go)
    
    * [练习代码](./lessons/lc5_longest_palindromic_substring.go)

9. [不同的子序列（Hard）](https://leetcode-cn.com/problems/distinct-subsequences/)

    * [练习代码](./lessons/lc115_distinct_subsequences.go)

10. [正则表达式匹配（Hard）](https://leetcode-cn.com/problems/regular-expression-matching/)

    * [练习代码](./lessons/lc10_regular_expression_matching.go)


-----------------------------------

图论(Graph theory)
------------------

#### 图的基本概念

* 图: 图可以标识为一个点集和一个边集 - Graph(V, E)

    + V - vertex: 点
        
        - 点的度数(degree)
        
            * 入度和出度: 一个点相连的入边/出边数量
            
    + E - edge: 边
    
        - 有向和无向
        
            * 带权图: 边的权值（长度）
       
#### 图的存储        

* 邻接矩阵

* 出边数组

* 邻接表

#### 图的深度优先遍历


#### 图的广度优先遍历


最短路
-------------

#### 单源最短路径问题
> 单源最短路径问题（Single Source Shortest Path, SSSP 问题）是说

* 给定一张有向图 `G = (V, E)`, `V` 是点集，`E` 是边集, `|V| = n, |E| = m`

* 节点以 [1, n] 之间的连续整数编号

* (x, y, z) 描述一条从 x 出发，到达 y，长度为 z 的有向边

* 设 1 号点为起点

* 求长度为 `n` 的数组 `dist`，其中 `dist[i]` 表示从起点 1 到节点 i 的最短路径的长度


#### Bellman-Ford 算法
> Bellman-ford 算法是基于动态规划和迭代思想的

1. 扫描所有边`(x, y, z)`, 若`dist[y] > dist[x] + z`, 则用 `dist[x] + z` 更新 `dist[y]`

2. 重复上述步骤，直到没有更新操作发生

> 若问题有解(图中没有负环)，Bellman-Ford "扫描所有边并更新" 的过程至多执行 n - 1 轮

* 原因: 一条最短路至多包含 n - 1 条边

> 时间复杂度: O(N * M)

> 可以把每一轮看作 DP 的一个阶段，第 i 轮至少已经求出了包含边数不超过 i 的最短路 


#### Dijkstra 算法
> Dijkstra 算法是基于贪心思想的，只适用于所有边的长度都是非负数的图

1. 初始化 `dist[1] = 0`, 其余节点的 `dist` 值为正无穷大

2. 找出一个未被标记的、`dist[x]` 最小的节点 `x`, 然后标记节点 `x`

3. 扫描节点 `x` 的所有出边 `(x, y, z)`, 若 `dist[y] > dist[x] + z` 则使用 `dist[x] + z` 更新 `dist[y]`

4. 重复上述 2~3 两个步骤, 直到所有节点都被标记

> 贪心思路: 在非负权图上，全局最小的 `dist` 值不可能再被其它节点更新
  因此可以不断取 `dist` 最小的点（每个点只被取一次），更新其它点
  
> 用二叉堆维护最小 `dist` 值可以做到 `O(m*log(n))` 的时间复杂度

> Dijkstra 算法的另一种理解方法: 优先队列 BFS

* 如果边权都是 1, 求最短路，可以用 BFS。BFS 每个点只需要访问一次，因为队列中点对应的路径长度满足“单调性”和“两段性”

* 如果边权是任意非负数，怎么保证每个点依然只需要扩展一次？

    + 优先队列 BFS ------ 先扩展最短的


#### Floyd 算法
> Floyd 算法可以在 O(N^3) 时间内求出图中每一对点之间的最短路径，本质上是动态规划算法

* dp[k, i, j] 表示经过编号不超过 k 的点为中继，从 i 到 j 的最短路

* 决策: 是否使用 k 这个中继点

    + dp[k, i, j] = min(dp[k - 1, i, j], dp[k - 1, i, k] + dp[k - 1, k, j])

* 可以省掉第一纬, 变成
    
    + d[i, j] = min(d[i, j], d[i, k] + d[k, j])
    
* 初态: d 为邻接矩阵(原始图中的边)

* 与 Bellman-Ford, Dijkstra 比较: `O(n^3) vs O(n^2 * m) vs O(n*m*log(n))`


最小生成树(Minimum Spanning Tree, MST)
-----------------------------------------------
给定一张边带权的无向图 `G = (V, E), n = |V|, m = |E`
由 V 中全部 n 个顶点和 E 中 n - 1 条边构成的无向连通图被称为 G 的一个生成树。
边的权值之和最小的生成树被称为无向图 G 的最小生成树

* 性质: 任意一颗最小生成树一定包含无向图中权值最小的边

* 推论: 把任何一个生成森林扩展为生成树，一定使用了图中剩余边中权值最小的

* 证明方法: 树上加一条边形成环 + 反证法

#### Kruskal 算法
> Kruskal 算法总是使用并查集维护无向图的最小生成森林

1. 建立并查集, 每个点各自构成一个集合

2. 把所有边按照权值从小到大排序，一次扫描每条边 `(x, y, z)`

3. 若 `x, y` 属于同一个集合(连通), 则忽略这条边，继续扫描下一条

4. 否则，合并 `x, y` 所在的集合，并把 `z` 累加到答案中

5. 所有边扫描完成后，第 4 步中处理过的边就构成最小生成树

> 时间复杂度为: O(m*log(m))


Rabin-Karp 算法
---------------
Rabin-Karp 是一种基于 Hash 的搞笑的字符串搜索算法

#### 问题
> 给定长度为 n 的字符串 s(文本串)，长度为 m 的字符串 t(模式串), 求 t 是否在 s 中出现过 (t 是否为 s 的子串)

#### 解法思路

1. 朴素: O(M*N)

2. Rabin-Karp 算法: O(M+n)

    - 思路: 计算 s 的每个长度为 m 的字串的 Hash 值（一个宽度为 m 的滑动窗口滑过 s）, 检查是否与 t 的 Hash 值相等




