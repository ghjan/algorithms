package inthashtable

import (
	"math"
	"fmt"
)

type EntryType int

const MAXTABLESIZE = 100000
const (
	_          EntryType = iota
	Legitimate
	Empty
	Deleted
)

/*Cell 散列表单元类型 */
type Cell struct {
	Data int       /* 存放元素 */
	Info EntryType /* 单元状态 */
}

//IntHashTable 散列表类型
type IntHashTable struct {
	TableSize int    /* 表的最大长度 */
	Cells     []Cell /* 存放散列单元数据的数组 */
}

func NextPrime(N int) int {
	/*从大于N的下一个奇数开始 */
	var p int
	if N%2 == 1 {
		p = N + 2
	} else {
		p = N + 1
	}
	var i int
	for p <= MAXTABLESIZE {
		for i = int(math.Sqrt(float64(p))); i > 2; i-- {
			if (p % i) == 0 {
				break //p不是素数
			}
		}

		if i == 2 {
			break //for正常结束，说明p是素数
		} else {
			p += 2 //否则试探下一个奇数
		}
	}
	return p
}

func (table IntHashTable) Hash(Key int) int {
	return Key % table.TableSize
}

func (table IntHashTable) Find(Key int) int {
	CNum := 0 //记录冲突次数

	NewPos := table.Hash(Key) /* 初始散列位置 */
	CurrentPos := NewPos      //当前散列位置
	/* 当该位置的单元非空，并且不是要找的元素时，发生冲突 */
	for CNum = 0; table.Cells[NewPos].Info != Empty && table.Cells[NewPos].Data != Key && CNum < table.TableSize; CNum++ {
		/* 统计1次冲突，并判断奇偶次 */
		CNum ++
		if CNum%2 == 1 { /* 奇数次冲突 */
			NewPos = CurrentPos + (CNum+1)*(CNum+1)/4 /* 增量为+[(CNum+1)/2]^2 */
			if NewPos >= table.TableSize {
				NewPos = NewPos % table.TableSize /* 调整为合法地址 */
			}
		} else { /* 偶数次冲突 */
			NewPos = CurrentPos - CNum*CNum/4 /* 增量为-(CNum/2)^2 */
			for NewPos < 0 {
				NewPos += table.TableSize /* 调整为合法地址 */
			}
		}
	}
	if CNum >= table.TableSize {
		return -1
	}
	/* 此时NewPos或者是Key的位置，或者是一个空单元的位置（表示找不到）*/
	return NewPos
}

func (table IntHashTable) Insert(Key int) int {
	Pos := table.Find(Key) /* 先检查Key是否已经存在 */
	if Pos < 0 {
		return Pos
	}
	if table.Cells[Pos].Info != Legitimate { /* 如果这个单元没有被占，说明Key可以插入在此 */
		table.Cells[Pos].Info = Legitimate
		table.Cells[Pos].Data = Key
		return Pos
	} else {
		fmt.Println("键值已存在")
		return Pos
	}
}

func (table IntHashTable) FindLinear(Key int) int {
	CNum := 0 //记录冲突次数

	NewPos := table.Hash(Key) /* 初始散列位置 */
	CurrentPos := NewPos      //当前散列位置
	/* 当该位置的单元非空，并且不是要找的元素时，发生冲突 */
	for CNum = 0; table.Cells[NewPos].Info != Empty && table.Cells[NewPos].Data != Key && CNum < table.TableSize; CNum++ {

		NewPos = (CurrentPos + CNum + 1) % table.TableSize
	}
	if CNum >= table.TableSize {
		return -1
	}
	/* 此时NewPos或者是Key的位置，或者是一个空单元的位置（表示找不到）*/
	return NewPos
}

func (table IntHashTable) InsertLinear(Key int) int {
	Pos := table.FindLinear(Key) /* 先检查Key是否已经存在 */
	if Pos < 0 {
		return Pos
	}
	if table.Cells[Pos].Info != Legitimate { /* 如果这个单元没有被占，说明Key可以插入在此 */
		table.Cells[Pos].Info = Legitimate
		table.Cells[Pos].Data = Key
		return Pos
	} else {
		fmt.Println("键值已存在")
		return Pos
	}
}

func CreateHashTable(TableSize int) IntHashTable {
	var H IntHashTable
	/* 保证散列表最大长度是素数 */
	H.TableSize = NextPrime(TableSize)
	/* 声明单元数组 */
	/* 初始化单元状态为“空单元” */
	H.Cells = make([]Cell, H.TableSize, H.TableSize)
	for i := 0; i < H.TableSize; i++ {
		H.Cells[i].Data = -1
		H.Cells[i].Info = Empty
	}

	return H
}
