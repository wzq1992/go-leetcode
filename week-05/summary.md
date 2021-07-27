本周作业
----------
1. [在 D 天内送达包裹的能力（Medium）](https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/)

    * [作业代码](./homework/lc1011_capacity_to_ship_packages_within_d_days.go)

2. [在线选举（Medium](https://leetcode-cn.com/problems/online-election/)

    * [作业代码](./homework/lc911_online_election.go)
    
3. [爱吃香蕉的珂珂（Medium）](https://leetcode-cn.com/problems/koko-eating-bananas/)

    * [作业代码](./homework/lc875_koko_eating_bananas.go)
    
4. [区间和的个数（选做）（Hard）](https://leetcode-cn.com/problems/count-of-range-sum/)

    * [作业代码](./homework/lc327_count_of_range_sum.go)


课上例题
----------
1. [寻找峰值（Medium）](https://leetcode-cn.com/problems/find-peak-element/)

    * [练习代码](./lessons/lc162_find_peak_element.go)

2. [猜数字大小（Easy）](https://leetcode-cn.com/problems/guess-number-higher-or-lower/)

    * [练习代码](./lessons/lc374_guess_number_higher_or_lower.go)

3. [分割数组的最大值（Hard）](https://leetcode-cn.com/problems/split-array-largest-sum/)

    * [练习代码](./lessons/lc410_split_array_largest_sum.go)

4. [制作 m 束花所需的最少天数（Medium）](https://leetcode-cn.com/problems/minimum-number-of-days-to-make-m-bouquets/)

    * [练习代码](./lessons/lc1482_minimum_number_of_days_to_make_m_bouquets.go)

5. [排序数组（Medium）](https://leetcode-cn.com/problems/sort-an-array/)

    * [练习代码](./lessons/lc912_sort_an_array.go)

6. [数组的相对排序（Easy）](https://leetcode-cn.com/problems/relative-sort-array/)

    * [练习代码](./lessons/lc1122_relative_sort_array.go)
    
7. [合并区间（Medium）](https://leetcode-cn.com/problems/merge-intervals/)

    * [练习代码 - 二维 Slice 排序](./lessons/lc56_merge_intervals_by_sort.go)
    * [练习代码 - 差分](./lessons/lc56_merge_intervals_by_difference.go)
    
8. [数组中的第 K 个最大元素（Medium）](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/)
    
    * [练习代码](./lessons/lc215_kth_largest_element_in_an_array.go)

9. [翻转对（Hard）](https://leetcode-cn.com/problems/reverse-pairs/)

    * [练习代码](./lessons/lc493_reverse_pairs.go)

10. [分发饼干（Easy）](https://leetcode-cn.com/problems/assign-cookies/)

    * [练习代码](./lessons/lc455_assign_cookies.go)

11. [买卖股票的最佳时机 II（Easy）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)

    * [练习代码](./lessons/lc122_best_time_to_buy_and_sell_stock_ii.go)

12. [跳跃游戏 II （Medium）](https://leetcode-cn.com/problems/jump-game-ii/)

    * [练习代码](./lessons/lc45_jump_game_ii.go)

13. [完成所有任务的最少初始能量（Hard）](https://leetcode-cn.com/problems/minimum-initial-energy-to-finish-tasks/)

    * [练习代码](./lessons/lc1665_minimum_initial_energy_to_finish_tasks.go)


-----------------------------------

三分查找
----------
> 二分与三分

* 二分：用于在单调函数上寻找特定值

* 三分：用于在单峰函数上寻找极大值

> 三分法用于单峰函数的极大值（或者单谷函数的极小值）。三分法也可用于求函数的局部极大/极小值

* 要求：函数是分段严格单调递增/递减的（不能出现一段平的情况）

> 以求单峰函数 f 的极大值为例，可在定义域 [l,r] 上任取两点 lmid, rmid

* 若 f(lmid) <= f(rmid) 则函数必然在 lmid 处单调递增，极值在 [lmid, r] 上

* 若 f(lmid) > f(rmid) 则函数必然在 rmid 处单调递减，极值在 [l, rmid] 上

> lmid, rmid 可去三等分点，也可以取 lmid 为二等分点，rmid 为 lmid 加一点偏移量


二分答案
----------
> 对一个最优化问题

* 求解：求一个最优解（最大值/最小值）

* 判定：给一个解，判断它是否合法（是否能够实现）

> 二分答案的本质是建立一个单调分段 0/1 函数。
  定义域为解空间（答案），值域为 0 或 1，在这个函数上二分查找分界点。
  
  
基于比较的排序算法
---------------

#### 初级排序算法，平均时间复杂度 O(N^2)

* 选择排序（Selection Sort）

```go
package main

// 该放哪个位置了
// 每次从未排序数据中心找最小值，放到已排序序列的末尾
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
```

* 插入排序（Insertion Sort）

```go
package main

// 这个数该放哪里了
// 从前到后一次考虑每个未排序数据，在已排序序列中找到合适位置插入
// func InsertionSort(arr []int) {
// 	for i := 1; i < len(arr); i++ {
// 		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
// 			arr[j], arr[j-1] = arr[j-1], arr[j]
// 		}
// 	}
// }

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		e := arr[i]
		var j int
		for j = i; j > 0 && arr[j-1] > e; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}
```

* 冒泡排序（Bubble Sort）

```go
package main

// 不断循环扫描，每次查看相邻元素，如果逆序则交换

```
  

#### 初级排序算法优化，平均时间复杂度 O(N log N)

* 堆排序（Heap Sort）

```go
package main

// 堆排序是对选择排序的优化，利用二叉堆高效的选出最小值
// 建立一个包含所有 N 个元素的二叉堆
// 重复 N 次从堆中取出最小值，即可得到有序序列
func heapSort(a []int, n int) {
	
}
``` 
  
* 希尔排序（Shell Sort）

```go
package main

// 希尔排序是对插入排序的优化，增量分组插入排序
// 希尔排序的时间复杂度取决于增量序列（步长序列）的选取
// 目前已知的最好序列可以做到 O(N^4/3) 或 O(N log^2N)
```
  
* 归并排序（Merge Sort）

```go
package main

// 归并排序是一个基于分治的算法
// 原问题：胞数组排序
// 子问题：把数组分成左右两半分别排序
// 然后合并

func mergeSort() {
	
}
```
  
* 快速排序（Quick Sort）

```go
package main

import (
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().Unix()))

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)

	return nums
}

// 快速排序也是一个基于分治的算法
// 从数组中选取中轴元素 pivot
// 然后分别对左右两边子数组进行快排

// 快排和归并排序具有相似性，但步骤顺序相反
// 归并排序：先排序左右子数组，然后合并两个有序数组
// 快速排序：先调配出左右子数组，然后对左右子数组分别进行排序
func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	pivot := partition(nums, l, r)
	quickSort(nums, l, pivot)
	quickSort(nums, pivot+1, r)
}

func partition(nums []int, l, r int) int {
	pivot := l + rnd.Int()%(r-l+1)
	pivotVal := nums[pivot]

	for l <= r {
		for nums[l] < pivotVal {
			l++
		}

		for nums[r] > pivotVal {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	return r
}

```


#### 排序的稳定性
> 对于序列中存在的若干大小相等的元素（重复元素）。
  如果排序前后它们的相对次序保持不变，就称排序算法是稳定的，
  否则就称排序算法是不稳定的
  
* 插入、冒泡、归并、计数、基数和桶排序是稳定的

* 选择、希尔、快速、堆排序是不稳定的


非比较排序
----------

#### 计数排序
> 计数排序要求输入的数据必须是有确定范围的整数。
  将输入的数据作为 key 存储在额外的数组中，然后依次把计数大于 1 的填充回原数组
  
* 时间复杂度：O(N + M), N 为元素个数，M 为数值范围

#### 桶排序
> 桶排序假设输入数据服从均匀分布，将数据分到有限数量的桶里，每个桶再分别排序（有可能使用别的排序算法，或是以递归的方式继续使用桶排序）

* 时间复杂度：O(N) ~ O(N^2)

#### 基数排序
> 基数排序把数据切割成一位位数字（0-9），从低位到高位对每一位分别进行计数排序

* 时间复杂度：O(NK), K 为数字位数


贪心算法(Greedy Algorithm)
-------------------------
> 贪心算法是一种

1. 在每一步选择中都采取在当前状态下最优决策(局部最优)

2. 并希望由此导致的最终结果也是全局最优

#### 贪心算法的证明

* 决策包容性

* 决策范围扩展
    
    在思考贪心算法时，有时候不容易直接证明局部最优决策的正确性，
    此时可以往后多扩展一步，有助于对当前的决策进行验证

* 邻项交换
    
    经常用于以某种顺序 "排序" 为贪心策略的证明，证明在任意局面下，
    任何局部的逆序改变都会造成整体结果变差
