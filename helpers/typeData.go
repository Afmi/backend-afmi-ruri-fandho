package helpers

type TypeData struct{}

func (h *TypeData) StrPtr(s string) *string {
	return &s
}
func (h *TypeData) IntPtr(s int) *int {
	return &s
}
