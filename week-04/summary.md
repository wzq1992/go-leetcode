本周作业
----------
1. [设计推特（Medium）](https://leetcode-cn.com/problems/design-twitter/)

    * [作业代码](./homework/lc355_design_twitter.go)

2. [数据流的中位数（选做）（Hard）](https://leetcode-cn.com/problems/find-median-from-data-stream/)

    * [作业代码](./homework/lc295_find_median_from_data_stream.go)
    
3. [寻找旋转排序数组中的最小值 II （Hard）](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/)

    * [作业代码](./homework/lc154_find_minimum_in_rotated_sorted_array_ii.go)


课上例题
----------
1. [最小基因变化（Medium）](https://leetcode-cn.com/problems/minimum-genetic-mutation/)

    * [练习代码](./lessons/lc433_minimum_genetic_mutation.go)

2. [矩阵中的最长递增路径（Hard）](https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix/)

    * [DFS 练习代码](./lessons/lc329_longest_increasing_path_in_a_matrix_dfs.go)
    
    * [BFS 练习代码](./lessons/lc329_longest_increasing_path_in_a_matrix_bfs.go)

3. [合并K个升序链表（Hard）](https://leetcode-cn.com/problems/merge-k-sorted-lists/)

    * [练习代码](./lessons/lc23_merge_k_sorted_lists.go)

4. [滑动窗口最大值（Hard）](https://leetcode-cn.com/problems/sliding-window-maximum/)

    * [练习代码](./lessons/lc239_sliding_window_maximum.go)

5. [二叉搜索树中的插入操作（Medium）](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)

    * [练习代码](./lessons/lc701_insert_into_a_binary_search_tree.go)
    
6. [面试题 04.06. 后继者（Medium）](https://leetcode-cn.com/problems/successor-lcci/)

    * [练习代码](./lessons/lc_successor_lcci.go)

7. [删除二叉搜索树中的节点（Medium）](https://leetcode-cn.com/problems/delete-node-in-a-bst/)

    * [练习代码](./lessons/lc450_delete_node_in_a_bst.go)

8. [二分查找（Easy）](https://leetcode-cn.com/problems/binary-search/)

    * [练习代码](./lessons/lc704_binary_search.go)
       
9. [在排序数组中查找元素的第一个和最后一个位置（Medium）](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/)

    * [练习代码](./lessons/lc34_find_first_and_last_position_of_element_in_sorted_array.go)

10. [x 的平方根（Easy）](https://leetcode-cn.com/problems/sqrtx/)

    * [练习代码](./lessons/lc69_sqrtx.go)
    
11. [寻找旋转排序数组中的最小值（Medium）](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)

    * [练习代码](./lessons/lc153_find_minimum_in_rotated_sorted_array.go)


-----------------------------------

堆
--------
> 堆(Heap) 是一种高效维护集合中最大或最小元素的数据结构

* 大根堆: 根节点最大的堆，用于维护和查询 max

* 小根堆: 根节点最小的堆，用于维护和查询 min

> 堆是一颗二叉树，并且满足堆性质（Heap property）

* 大根堆任意节点的 key >= 所有子节点 key

* 小根堆任意节点的 key <= 所有子节点 key


二叉堆（Binary Heap）
-------------------
二叉堆是堆得一种简易实现，本质上是一颗满足堆性质的完全二叉树

> 常见操作

* 建堆（build）: O(N)

* 查询最值（get max/min）: O(1)

* 插入（insert）: O(log N)

* 取出最值（delete max/min）: O(log N)


二叉堆的实现
-----------
二叉堆一般使用一个一维数组来存储，利用完全二叉树的节点编号特性

```go
// source code: /usr/local/go/src/container/heap/heap_test.go

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
```

优先队列
----------
二叉堆是优先队列（Priority Queue）一种简单、常见的实现，但不是最优实现

```go
// source code: /usr/local/go/src/container/heap/example_pq_test.go

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
```

二叉搜索树（Binary Search Tree）
-----------------------------
> 中序遍历是有序的

#### BST 建立

* 两个保护节点: 仅由这两个节点构成的 BST 就是一颗初始的空 BST
    
    + 正无穷
    
    + 负无穷

#### BST 检索

#### BST 插入
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val:val}
    }

    if val < root.Val {
        root.Left = insertIntoBST(root.Left, val)
    } else {
        root.Right = insertIntoBST(root.Right, val)
    }

    return root
}
```

#### BST 求前驱、后继

> 求后继
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    return findSucc(root, p.Val)
}

func findSucc(root *TreeNode, val int) *TreeNode {
    curr := root
    var ans *TreeNode

    for curr != nil {
        if curr.Val > val && (ans == nil || ans.Val > curr.Val) {
            ans = curr
        }

        if curr.Val == val && curr.Right != nil {
            curr = curr.Right

            for curr.Left != nil {
                curr = curr.Left
            }

            return curr
        }

        if val < curr.Val {
            curr = curr.Left
        } else {
            curr = curr.Right
        }
    }

    return ans
}
```

#### BST 删除

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return root
    }

    if root.Val == key {
        if root.Left == nil {
            return root.Right
        }

        if root.Right == nil {
            return root.Left
        }

        next := root.Right
        // 找后继 右子树一路向左
        for next.Left != nil {
            next = next.Left 
        }

        root.Right = deleteNode(root.Right, next.Val)
        root.Val = next.Val

        return root
    }

    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else {
        root.Right = deleteNode(root.Right, key)
    }

    return root
}
```


二分查找
-------

#### 二分查找的前提

* 目标函数具有单调性

* 存在上下界

* 能够通过索引访问

#### 代码模板

```go
left, right := 0, n - 1
for lef <= right {
	mid := left + (right - left) / 2
	
	if arr[mid] == target {
		return mid
	}
	
	if arr[mid] < target {
		left = mid + 1
	} else {
		right = mid
	}
}
```

#### lower_bound
> 在单调递增数组里，查找第一个 >= 指定数的下标，不存在返回数组的长度

#### upper_bound
> 在单调递增数组里，查找第一个 > 指定数的下标[找后继]，不存在返回数组的长度

#### lower_bound 与 upper_bound 区别
> 给定的 target 不一定存在于数组中, arr[mid] 即使不等于 target, 也可能就是最后的答案，不能随便排除在外

* (1.1) + (1.2): 最严谨的划分，一侧包含，一侧不包含，终止与 left == right

* (2): 双侧都不包含，用 ans 维护答案， 终止于 left > right

* (3): 双侧都包含，终止于 left + 1 = right, 最后再检查答案

#### 二分模板(1.1)

```go
left, right := 0, n

for left < right {
	mid := left + (right - left) / 2
	
	// upper_bound 就是 arr[mid] > target
	if arr[mid] >= target {
		right = mid
	} else {
		left = mid + 1
	}
}

return right
```

#### 二分模板(1.2)

```go
// 查找最后䘝 <= target 的数，不存在返回 -1
left, right := -1, n - 1

for left < right {
	mid := (left + right + 1) / 2
	
	if arr[mid] <= target {
		left = mid
	} else {
		right = mid - 1
	}
}

return right
```

#### 解题步骤

1. 条件

2. 范围

3. 补 +1