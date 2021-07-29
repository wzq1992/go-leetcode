本周作业
----------
1. [爬楼梯（Easy）](https://leetcode-cn.com/problems/climbing-stairs/)

    * [作业代码](./homework/lc70_climbing_stairs.go)

2. [三角形最小路径和（Medium）](https://leetcode-cn.com/problems/triangle/)

    * [作业代码](./homework/lc120_triangle.go)


课上例题
----------
1. [零钱兑换（Medium）](https://leetcode-cn.com/problems/coin-change/)

    * [练习代码](./lessons/lc322_coin_change.go)

2. [不同路径 II （Medium）](https://leetcode-cn.com/problems/unique-paths-ii/)

    * [练习代码](./lessons/lc63_unique_paths_ii.go)

3. [最长公共子序列（Medium）](https://leetcode-cn.com/problems/longest-common-subsequence/)

    * [练习代码](./lessons/lc1143_longest_common_subsequence.go)

4. [最长递增子序列（Medium）](https://leetcode-cn.com/problems/longest-increasing-subsequence/)

    * [练习代码](./lessons/lc300_longest_increasing_subsequence.go)

5. [最大子序和（Easy）](https://leetcode-cn.com/problems/maximum-subarray/)

    * [练习代码](./lessons/lc53_maximum_subarray.go)
    
6. [乘积最大子数组（Medium）](https://leetcode-cn.com/problems/maximum-product-subarray/)
    
    * [练习代码](./lessons/lc152_maximum_product_subarray.go)

7. [买卖股票的最佳时机（Easy）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)

    * [练习代码](./lessons/lc121_best_time_to_buy_and_sell_stock.go)

8. [买卖股票的最佳时机 II （Easy）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)

    * [练习代码-贪心](./lessons/lc122_best_time_to_buy_and_sell_stock_ii_01.go)
    
    * [练习代码-动态规划](./lessons/lc122_best_time_to_buy_and_sell_stock_ii_02.go)

9. [买卖股票的最佳时机 III （Easy）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)

    * [练习代码](./lessons/lc123_best_time_to_buy_and_sell_stock_iii.go)

10. [买卖股票的最佳时机 IV （Hard）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)

    * [练习代码](./lessons/lc188_best_time_to_buy_and_sell_stock_iv.go)

11. [买卖股票的最佳时机含手续费（Medium）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)

    * [练习代码](./lessons/lc714_best_time_to_buy_and_sell_stock_with_transaction_fee.go)

12. [最佳买卖股票时机含冷冻期（Medium）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)

    * [练习代码](./lessons/lc309_best_time_to_buy_and_sell_stock_with_cooldown.go)

13. [打家劫舍（Medium）](https://leetcode-cn.com/problems/house-robber/)

    * [练习代码](./lessons/lc198_house_robber.go)

14. [打家劫舍 II- 环形 DP （Medium）](https://leetcode-cn.com/problems/house-robber-ii/)

    * [练习代码](./lessons/lc213_house_robber_ii.go)

15. [编辑距离（重点题）](https://leetcode-cn.com/problems/edit-distance/)

    * [练习代码](./lessons/lc72_edit_distance.go)

16. [零钱兑换 II （Medium）](https://leetcode-cn.com/problems/coin-change-2/)

    * [练习代码](./lessons/lc518_coin_change_2.go)

17. [分割等和子集（Medium）](https://leetcode-cn.com/problems/partition-equal-subset-sum/)

    * [练习代码](./lessons/lc416_partition_equal_subset_sum.go)


-----------------------------------


动态规划(DP, Dynamic programming)
---------------------------------
动态规划是一种对问题的状态空间进行分阶段、有顺序、不重复决策性遍历的算法

#### 动态规划的关键与前提

* 重叠子问题: 与递归、分治等一样，要具有同类子问题，用若干纬状态表示

* 最优子结构: 状态对应着一个最优化目标，并且最优化目标之间具有推导关系

* 无后效性: 问题的状态空间是一张有向无环图（可按一定的顺序遍历求解）


#### 动态规划的实现
> 一般采用递推的方式实现，也可以写成递归或搜索的形式，因为每个状态只遍历一次，
  也被称为记忆化搜索

#### 动态规划三要素

* 阶段

* 状态

* 决策

#### 一道动态规划题目的标准题解

* 状态转义方程

#### 动态规划解题步骤

1. 人力模拟、蛮力搜索

2. 定义状态

3. 确定最优子结构

4. 写出状态转移方程

5. 确定边界、目标，递推实现

#### 动态规划打印方案

> 动态规划题目打印方案的原则：记录转移路径 + 递归输出

> 动态规划选取“代表”，维护了一个最优子结构，如果记录每个最优子结构的详细方案，
  时间复杂度会上升
  
> 正确做法：记录每个 `f[i,j]` 是从哪里转移过来的 `(f[i - 1, j], f[i, j - 1])`
  还是 `f[i -1, j - 1]` 整个动规完成，求出 `f[m, n]` 后，再从 `(m, n)` 开始递归
  复原最优路径
  


