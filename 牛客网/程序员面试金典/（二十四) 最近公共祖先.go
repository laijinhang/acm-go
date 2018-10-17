/*
题目描述
有一棵无穷大的满二叉树，其结点按根结点一层一层地从左往右依次编号，根结点编号为1。现在有两个结点a，b。请设计一个算法，求出a和b点的最近公共祖先的编号。

给定两个int a,b。为给定结点的编号。请返回a和b的最近公共祖先的编号。注意这里结点本身也可认为是其祖先。

测试样例：
2，3
返回：1

分析：根据题意可推出
1）存在一个编号a，如果a!=1，则，它的最近祖先的编号为 a/2 向下取整，为一个整数。
2）假设a <=  b，存在以下几种情况
1、a = b      ->     a
2、b/2向下取整 n次 必然会等于a或者小于a，等于a则推出最近公共祖先为a，小于a则交换a，b值，直到a=b，即求出了最近公共祖先

开始编写代码：

 */
 package main

import "fmt"

func getLCA(a, b int) int {
	for a != b {
		if a > b {
			a, b = b, a
		}
		b /= 2
	}
	return a
}

func main() {
	fmt.Println(getLCA(2, 3))
}