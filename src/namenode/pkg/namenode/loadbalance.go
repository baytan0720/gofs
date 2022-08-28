package namenode

import "sort"

type load struct {
	id   int
	load float32
}

type LoadSlice []load

func (s LoadSlice) Len() int { return len(s) }

func (s LoadSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s LoadSlice) Less(i, j int) bool { return s[i].load > s[j].load }

//计算负载并从低到高排序
func (nn *NameNode) calLoad() []int {
	all := make([]load, 0, nn.DataNodeNum)
	for i, v := range nn.DataNodeList {
		if v == nil {
			continue
		}
		vload := v.load.Percpu*0.3 + v.load.Permem*0.2 + v.load.Perdisk*0.5
		if vload > float32(nn.MaxLoad) {
			continue
		}
		all = append(all, load{id: i, load: vload})
	}
	if len(all) < 3 {
		return nil
	}
	if len(all) > 3 {
		sort.Sort(LoadSlice(all))
	}
	res := make([]int, 3)
	for i := 0; i < 3; i++ {
		res[i] = all[i].id
	}
	return res
}
