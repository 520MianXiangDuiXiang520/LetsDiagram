package nottable

// FullCanvas 表示完整的一条 Canvas 记录
type FullCanvas struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Author uint   `json:"author"`
	Cover  string `json:"cover"`
	Data   string `json:"data"`
}
