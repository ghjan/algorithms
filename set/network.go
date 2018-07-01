package set

type IntSet []int //保存parent的index

//FindRoot 寻找某个节点的根节点
// func (S IntSet) FindRoot(X int) int {
// 	//默认集合元素全部初始化为-1, root节点保存-1*树元素个数（按秩归并-比规模)
// 	for ; S[X] >= 0; X = S[X] {
// 	}
// 	return X
// }

//FindRoot2 寻找某个节点的根节点（优化：路径压缩）
func (S IntSet) FindRoot2(X int) int {
	if S[X] < 0 { // 找到集合的根
		return X
	} else {
		S[X] = S.FindRoot2(S[X])
		return S[X]
	}
}

//Union 集合合并 按秩归并(比规模)
func (S IntSet) Union(Root1, Root2 int) {
	//假设root1和root2分别是两个不同集合的根节点
	if S[Root2] < S[Root1] {
		S[Root2] += S[Root1]
		S[Root1] = Root2
	} else {
		S[Root1] += S[Root2]
		S[Root2] = Root1
	}
}

//Initialization 初始化
func Initialization(n int) IntSet {
	S := make([]int, n, n)
	for i := 0; i < n; i++ {
		S[i] = -1
	}
	return S
}

//InputConnection 两个节点相连
func (S IntSet) InputConnection(u, v int) {
	Root1 := S.FindRoot2(u - 1)
	Root2 := S.FindRoot2(v - 1)
	if Root1 != Root2 {
		S.Union(Root1, Root2)
	}
}

//CheckConnection 检查两个节点是否相连
func (S IntSet) CheckConnection(u, v int) bool{
	Root1 := S.FindRoot2(u - 1)
	Root2 := S.FindRoot2(v - 1)
	return Root1 == Root2
}

//CheckNetwork 检查网络有几个根节点（也就是有几个独立的component）
func (S IntSet) CheckNetwork(n int) int {
	counter := 0
	for i := 0; i < n; i++ {
		if S[i] < 0 {
			counter++
		}
	}

	return counter
}
