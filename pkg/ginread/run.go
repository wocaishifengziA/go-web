package ginread

func Run() {

}

// 定义树的节点
type Node struct {
	// 1. 保存当前字符
	C byte
	// 2. 若干个子节点
	Children []*Node
	// 3. 标识节点是否是一个单词最后的一个字符
	isWord bool
}

// 给子节点 列表添加新的节点
func (n *Node) Add(c byte) {
	// 构造节点
	node := Node{
		C: c,
	}
	// 定义的 slice 类型的，不用初始化，也可以正常使用append方法
	n.Children = append(n.Children, &node)
}

// 返回子节点 列表 长度
func (n *Node) Len() int {
	return len(n.Children)
}

// 设置节点的 isword
func (n *Node) SetIsword(flag bool) {
	n.isWord = flag
}

// 整棵树的根节点
type Trie struct {
	// 字典树的根节点
	Root *Node
}


// 随便初始化根节点为 "/" 吧
func (t *Trie) Init() *Trie {
	t.Root = &Node{
		'/',
		make([]*Node, 0),
		false,
	}
	return t
}


// 外部调用初始化一个 字典树
func New() *Trie { return new(Trie).Init() }


// 添加单词
// 时间复杂度：
func (t *Trie) Add(word string) {
	// 转换为 字符
	w := []byte(word)
	// 找到每一个节点
	cur := t.Root
	// 遍历 word的每一个字符
	for _, c := range w {
		// 1. 查找子节点中是否包含当前字符
		index := t.containsChar(c, cur.Children)
		if index == -1 {
			// 添加新节点
			cur.Add(c)
			// 更新cur 指向它
			index = cur.Len() - 1
		}
		// cur 指向下一个子节点
		cur = cur.Children[index]
	}
	// 将最后一个节点字符的 尾设置为 true
	cur.SetIsword(true)
}


// 查找子节点中是否包含指定字符： 私有方法
func (t *Trie) containsChar(c byte, childrens []*Node) int {
	// 线性遍历查找
	for i, node := range childrens {
		// 查找到既返回索引
		if node.C ==  c {
			return i
		}
	}
	return -1
}


// 判断是否包含指定的单词
// 不仅能找到单词，并且单词的结尾必须是 isWord = true 才说明有这个单词
func (t *Trie) Contains(word string) bool {
	// 转换为 字符
	w := []byte(word)
	// 找到每一个节点
	cur := t.Root
	// 遍历 word的每一个字符
	for _, c := range w {
		// 1. 查找子节点中是否包含当前字符
		index := t.containsChar(c, cur.Children)
		if index == -1 {
			return false
		}
		// cur 指向下一个子节点
		cur = cur.Children[index]
	}
	// 判断最后一个单词是否是 单词末尾 cur.isWord == true,
	// 直接缩写为 cur.isWord
	return cur.isWord
}
