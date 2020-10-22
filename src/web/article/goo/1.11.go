package main

// 如何判断两个单链表是否交叉
func main() {
	// hash法，先遍历head1，放到map里（地址），然后遍历head2，是否有重叠部分

	// 首尾相接法，即head1的尾节链接到head2的头指针，然后检测是否有环

	// 尾节点法，如果head1和head2有相同的尾部节点，那么head1和head2有交叉
	// 遍历可获取head1和head2的长度，长链表先走l1-l2步，之后head1和head2每次各走一步，相遇的第一点即为相交的第一个点
}
