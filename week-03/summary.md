本周作业
----------


课上例题
----------


-----------------------------------

二叉树
--------
Binary Tree

* 完全二叉树

* 满二叉树

* 树节点定义

    ```go
    type TreeNode struct {
        val int
        left, right *TreeNode
    }
    ```

* 二叉树的遍历

    + 前序遍历(Pre-Order): 根->左->右
    
    + 中序遍历(In-Order): 左->根->右
    
    + 后序遍历(Post-Order): 左->右->根
    
    * 层序遍历: 广度优先
