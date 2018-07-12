package set

//IntSet 并查集
type IntSet []int //保存parent的index

//FindRoot 寻找某个节点的根节点（优化：路径压缩）
func (S IntSet) FindRoot(X int) int {
	//默认集合元素全部初始化为-1
	if S[X] < 0 { // 找到集合的根
		return X
	} else {
		S[X] = S.FindRoot(S[X])
		return S[X]
	}
}

//Union 集合合并 保证小集合并入大集合 按秩归并(比规模)
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
//返回： 	true 表示uv原先没有连接，合并两个联通集
//			false 表示原先已经在同一个联通集
func (S IntSet) InputConnection(u, v int) bool {
	//检查连接u,v的边是否在现有的子集中构成回路
	Root1 := S.FindRoot(u - 1)
	Root2 := S.FindRoot(v - 1)
	isConnected := Root1 == Root2
	if !isConnected {
		// 否则该边可以被收集，同时将u和v并入同一连通集
		S.Union(Root1, Root2)
	}
	return !isConnected
}

//CheckCycle 检查是否有回路(u,v已经有通路），没有回路的情况下面合并两个集合
//说明：和InputConnection函数重复了
//返回： 	true 表示uv原先没有连接，合并两个联通集
//			false 表示原先已经在同一个联通集
func (S IntSet) CheckCycle(u, v int) bool {
	return S.InputConnection(u, v)
}

//CheckConnection 检查两个节点是否相连
func (S IntSet) CheckConnection(u, v int) (bool) {
	Root1 := S.FindRoot(u - 1)
	Root2 := S.FindRoot(v - 1)
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
