# 快速排序

快速排序是冒泡排序的优化版，同属于**交换类排序**。它使用了 “分治法” 的思想，将集合分割为相似子集，再对子集进行递归排序，最后将子集的排序结果组合即可。

## 排序过程

```shell
quick_sort $ go run main.go
[UNSORTED]:      [27 38 12 39 27 16]
[DEBUG low]:     []
[DEBUG high]:    [16]
[DEBUG low]:     [27]
[DEBUG high]:    [39]
[DEBUG low]:     [12 16]
[DEBUG high]:    [27 38 39]
[SORTED]:        [12 16 27 27 38 39]
```

## 排序效果

![4](http://p7f8yck57.bkt.clouddn.com/2018-06-14-042424.gif)

## 复杂度

### 时间复杂度

最好情况：递归求证，**O( NlogN )**

最坏情况：退化为冒泡排序 **O( N^2 )**

平均复杂度：**O(NlogN)**

### 空间复杂度

由于递归借助栈空间进行 **O(longN)** 次调用，本实现的空间复杂度为  **O(longN)**

### 稳定性

本实现的原地分区在比较值的时不区分值是否重复，是不稳定的。

## 使用场景

适用于大规模无序数据排序。



## 总结

快排的关键是“分治”的思想，对应到代码实现即递归。需注意递归退出 `return` 的条件，本例是子集中只有 1 个元素时视为排序完毕，递归结束。快排还有其他两种优化实现，可参考讨论：[how-to-optimize-quicksort](https://stackoverflow.com/questions/12454866/how-to-optimize-quicksort)



