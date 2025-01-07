### 按照时间复杂度分类:
- $O(N^2)$:
  - 冒泡排序（Bubble sort）
  - 选择排序（Selection sort）
  - 插入排序（Insertion sort）

- $O(Nlog_2N)$: 
  - 合并排序（Merge sort）
  - 快速排序（Quick sort）

- $O(N)$:
  - 桶排序（Bucket sort）
  - 计数排序（Counting sort）
  - 基数排序（Radix sort）
  
### 方法：
- 选择排序：在i ~ n-1范围内，找到最小值，并将其于i位置元素进行交换；然后在i+1 ~ n-1范围上继续之前的操作
- 冒泡排序：在0 ~ i范围内，依次两两比较使偏大的数交换到后面的位置；然后在0 ~ i-1范围上继续之前的操作
- 插入排序：在0 ~ i范围已经是有序的了，将i+1位置的数据按照从右向左依次比较的方式，选择合适位置插入；然后使用相同的操作插入i + 2位置的数据

### 参考资料：
- https://en.wikipedia.org/wiki/Sorting_algorithm
- https://www.programiz.com/dsa/sorting-algorithm