本周作业
----------
1. [冗余连接（Medium）](https://leetcode-cn.com/problems/redundant-connection/)

    * [作业代码](homework/lc684_redundant_connection.go)

2. [岛屿数量（Medium）](https://leetcode-cn.com/problems/number-of-islands/)

    * [作业代码](homework/lc200_number_of_islands.go)


课上例题
----------
1. [满足不等式的最大值（Hard）](https://leetcode-cn.com/problems/max-value-of-equation/)

    * [练习代码](./lessons/lc1499_max_value_of_equation.go)

2. [环形子数组的最大和（Medium）](https://leetcode-cn.com/problems/maximum-sum-circular-subarray/)

    * [练习代码](./lessons/lc918_maximum_sum_circular_subarray.go)

3. [戳气球（Hard）](https://leetcode-cn.com/problems/burst-balloons/)

    * [练习代码](./lessons/lc312_burst_balloons.go)

4. [合并石头的最低成本（Hard）](https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/)


5. [打家劫舍 III （Medium）](https://leetcode-cn.com/problems/house-robber-iii/)

    * [练习代码](./lessons/lc337_house_robber_iii.go)

6. [实现 Trie (前缀树) （Medium）](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)

    * [练习代码](./lessons/lc208_implement_trie_prefix_tree.go)

7. [单词搜索 II （Hard）](https://leetcode-cn.com/problems/word-search-ii/)

    * [练习代码](./lessons/lc212_word_search_ii.go)

8. [省份数量（Medium）](https://leetcode-cn.com/problems/number-of-provinces/)

    * [练习代码](./lessons/lc547_number_of_provinces.go)

9. [被围绕的区域（Medium）](https://leetcode-cn.com/problems/surrounded-regions/)

    * [练习代码](./lessons/lc130_surrounded_regions.go)


-----------------------------------


动态规划的优化
---------------

#### 动态规划的转移优化思想

* 分离 `i` 和 `j`。与 `i` 有关的式子放一边，不要互相干扰。

* 观察内层循环变量 `j` 的取值范围随着外层循环变量的变化情况

> 在动态规划中经常遇到类似的式子，`i` 是状态变量，`j` 是决策变量

* 分离状态变量和决策变量。当循环多于两重时，关注最里面的两重循环，把外层循环看做定值

* 对于一个状态变量，决策变量的取值范围称为“决策候选集合”，观察这个集合随着状态变量的变化情况

> 一旦发现冗余，或者有更高效维护“候选集合”的数据结构，就可以省去一层循环扫描


字典树(Trie)
-----------------
一种由“节点”和“带有字符的边”构成的属性结构。典型应用是用于统计和排序大量的字符串
（但不仅限于字符串），经常被搜索引擎系统用于文本词频统计。它的优点是：最大限度地
减少无谓的字符串比较，查询效率比哈希表高

#### 基本性质

* 节点本身不保存完整单词

* 从根及节点到某一个节点，路径上经过的字符串连接起来，为该节点对应的单词

* 每个节点出发的所有边代表的字符都不相同

* 节点用于存储单词的额外信息（例如词频）

#### 内部实现

1. 字符集数组法（简单）

    > 每个节点保存一个长度固定为字符集大小(例如 26)的数组，以字符为下标，保存指向的节点。
      空间复杂度为 O(节点数 * 字符集大小)，查询的时间复杂度为 O(单词长度)。
      适用于较小字符集，或者单词短、分布稠密的字典

2. 字符集映射法（优化）

    > 把每个节点上的字符集数组改为一个映射(词频统计：hash map, 排序 ordered map)。
      空间复杂度为 O(文本字符总数)，查询的时间复杂度为 O(单词长度)，但常数稍大一些。
      适用性更广


并查集(Disjoint sets)
---------------------

#### 基本用途

* 处理不相交集合(disjoint sets) 的合并和查询问题

* 处理分组问题

* 维护无序二元关系

#### 基本操作

* MakeSet(s): 建立一个新的并查集，其中包含 s 个集合，每个集合里只有一个元素

* UnionSet(x, y): 把元素 x 和 y 所在的集合合并，要求 x 和 y 所在集合不相交，如果相交则无需合并

* Find(x): 找到元素 x 坐在的集合的代表。该操作也可以用于判断两个元素是否位于同一个集合，只要将它们各自的代表比较一下就可以了。

#### 内部实现

* 每个集合是一个树形结构

* 每个节点只需要保存一个值：它的父节点

* 最简单的实现是只用一个 int 数组 `fa`, `fa[x]` 表示编号为 x 的节点的父节点，根节点的 fa 等于它自己

```go
package lessons

var fa []int

func DisjointSet(n int) {
	fa = make([]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
	}
}

func Find(x int) int {
	if x == fa[x] {
		return  x
	}
	
	fa[x] = Find(fa[x])
	return fa[x]
}

func unionSet(x, y int) {
	x, y = Find(x), Find(y)
	
	if x != y {
		fa[x] = y
	}
}

```