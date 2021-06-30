本周作业
----------
1. [LRU 缓存机制（Medium）](https://leetcode-cn.com/problems/lru-cache/)

    * [作业代码](./homework/lc146_lru_cache.go)
    
    * [测试用例](./homework/lc146_lru_cache_test.go)
    
2. [子域名访问计数（Easy）](https://leetcode-cn.com/problems/subdomain-visit-count/) 

    * [作业代码](./homework/lc811_subdomain_visit_count.go)
    
    * [测试用例](./homework/lc811_subdomain_visit_count_test.go)
    
3. [数组的度（Easy）](https://leetcode-cn.com/problems/degree-of-an-array/)

    * [作业代码](./homework/lc697_degree_of_an_array.go)
    
    * [测试用例](./homework/lc697_degree_of_an_array_test.go)
    
4. [元素和为目标值的子矩阵数量（Hard）](https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target/)

    * [作业代码](./homework/lc1074_number_of_submatrices_that_sum_to_target.go)
    
    * [测试用例](./homework/lc1074_number_of_submatrices_that_sum_to_target_test.go)

5. [合并K 个升序链表（Hard）](https://leetcode-cn.com/problems/merge-k-sorted-lists/) (要求：用分治实现，不要用堆)

    * [作业代码](./homework/lc23_merge_k_sorted_lists.go)
    
    * [测试用例](./homework/lc23_merge_k_sorted_lists_test.go)
    

课上例题
----------

1. [两数之和（Easy）](https://leetcode-cn.com/problems/two-sum/description/)

    * [练习代码](./lessons/lc1_two_sum.go)

2. [模拟行走机器人（Easy）](https://leetcode-cn.com/problems/walking-robot-simulation/)

    * [练习代码](./lessons/lc874_walking_robot_simulation.go)

3. [字母异位词分组（Medium）](https://leetcode-cn.com/problems/group-anagrams/)

    * [练习代码](./lessons/lc49_group_anagrams.go)

4. [串联所有单词的子串（Hard）](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/)

    * [练习代码](./lessons/lc30_substring_with_concatenation_of_all_words.go)

5. [子集（Medium）](https://leetcode-cn.com/problems/subsets/)

    * [练习代码](./lessons/lc78_subsets.go)

6. [组合（Medium）](https://leetcode-cn.com/problems/combinations/)

    * [练习代码](./lessons/lc77_combinations.go)

7. [全排列（Medium）](https://leetcode-cn.com/problems/permutations/)

    * [练习代码](./lessons/lc46_permutations.go)


-----------------------------------

哈希表
--------
Hash table, 散列表

组成: 

    * 数据结构: 链表、数组
    
    * Hash 函数

哈希碰撞:
    
    * 解决方案: 开散列

应用:
    
    * 电话号码簿
    
    * 用户信息表
    
    * 缓存（LRU Cache）
    
    * 键值对存储（Redis）

均匀分布的方法:

    * 模一个质数


集合与映射
---------

* 有序集合

* 无序集合


递归
-----------

* 函数自身调用自己

* 典型案例：阶乘
    
    ```go
    func factorial(n int) int {
        if n == 1 {
      	    return 1
        }
      
        return n * factorial(n - 1) 
    }
    ```
    
 * 关键点：子问题、边界、还原现场
 
 * 递归形式：指数型、排列型、组合型
 
 
 分治
 ---------
 把原问题划分成若干个同类子问题，分别解决后，合并
 
 * 子问题划分标准：不重不漏
 
 
 
 
 
 