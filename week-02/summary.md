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


-----------------------------------

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
 
 
 
 
 
 