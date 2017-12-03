package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkNotebook
//  A tabbed notebook container.
//
// Functions:
// GtkWidget * 	gtk_notebook_new ()
// gint 	gtk_notebook_append_page ()
// gint 	gtk_notebook_append_page_menu ()
// gint 	gtk_notebook_prepend_page ()
// gint 	gtk_notebook_prepend_page_menu ()
// gint 	gtk_notebook_insert_page ()
// gint 	gtk_notebook_insert_page_menu ()
// void 	gtk_notebook_remove_page ()
// - void 	gtk_notebook_detach_tab () 						-
// gint 	gtk_notebook_page_num ()
// void 	gtk_notebook_next_page ()
// void 	gtk_notebook_prev_page ()
// void 	gtk_notebook_reorder_child ()
// void 	gtk_notebook_set_tab_pos ()
// void 	gtk_notebook_set_show_tabs ()
// void 	gtk_notebook_set_show_border ()
// void 	gtk_notebook_set_scrollable ()
// void 	gtk_notebook_popup_enable ()
// void 	gtk_notebook_popup_disable ()
// gint 	gtk_notebook_get_current_page ()
// GtkWidget * 	gtk_notebook_get_menu_label ()
// GtkWidget * 	gtk_notebook_get_nth_page ()
// gint 	gtk_notebook_get_n_pages ()
// GtkWidget * 	gtk_notebook_get_tab_label ()
// void 	gtk_notebook_set_menu_label ()
// void 	gtk_notebook_set_menu_label_text ()
// void 	gtk_notebook_set_tab_label ()
// void 	gtk_notebook_set_tab_label_text ()
// void 	gtk_notebook_set_tab_reorderable ()
// void 	gtk_notebook_set_tab_detachable ()
// const gchar * 	gtk_notebook_get_menu_label_text ()
// gboolean 	gtk_notebook_get_scrollable ()
// gboolean 	gtk_notebook_get_show_border ()
// gboolean 	gtk_notebook_get_show_tabs ()
// const gchar * 	gtk_notebook_get_tab_label_text ()
// GtkPositionType 	gtk_notebook_get_tab_pos ()
// gboolean 	gtk_notebook_get_tab_reorderable ()
// gboolean 	gtk_notebook_get_tab_detachable ()
// - guint16 	gtk_notebook_get_tab_hborder () 			-
// - guint16 	gtk_notebook_get_tab_vborder () 			-
// void 	gtk_notebook_set_current_page ()
// void 	gtk_notebook_set_group_name ()
// const gchar * 	gtk_notebook_get_group_name ()
// - void 	gtk_notebook_set_action_widget () 				-
// - GtkWidget * 	gtk_notebook_get_action_widget ()		-
//-----------------------------------------------------------------------

type Notebook struct {
	Container
}

func NewNotebook() *Notebook {
	return &Notebook{Container{Widget{
		C.gtk_notebook_new()}}}
}

func (v *Notebook) AppendPage(child IWidget, tab_label IWidget) int {
	return int(C.gtk_notebook_append_page(NOTEBOOK(v), ToNative(child), ToNative(tab_label)))
}

func (v *Notebook) AppendPageMenu(child IWidget, tab_label IWidget, menu_label IWidget) int {
	return int(C.gtk_notebook_append_page_menu(NOTEBOOK(v), ToNative(child), ToNative(tab_label), ToNative(menu_label)))
}

func (v *Notebook) PrependPage(child IWidget, tab_label IWidget) int {
	return int(C.gtk_notebook_prepend_page(NOTEBOOK(v), ToNative(child), ToNative(tab_label)))
}

func (v *Notebook) PrependPageMenu(child IWidget, tab_label IWidget, menu_label IWidget) int {
	return int(C.gtk_notebook_prepend_page_menu(NOTEBOOK(v), ToNative(child), ToNative(tab_label), ToNative(menu_label)))
}

func (v *Notebook) InsertPage(child IWidget, tab_label IWidget, position int) int {
	return int(C.gtk_notebook_insert_page(NOTEBOOK(v), ToNative(child), ToNative(tab_label), gint(position)))
}

func (v *Notebook) InsertPageMenu(child IWidget, tab_label IWidget, menu_label IWidget, position int) int {
	return int(C.gtk_notebook_insert_page_menu(NOTEBOOK(v), ToNative(child), ToNative(tab_label), ToNative(menu_label), gint(position)))
}

func (v *Notebook) RemovePage(child IWidget, page_num int) {
	C.gtk_notebook_remove_page(NOTEBOOK(v), gint(page_num))
}

func (v *Notebook) PageNum(child IWidget) int {
	return int(C.gtk_notebook_page_num(NOTEBOOK(v), ToNative(child)))
}

func (v *Notebook) NextPage() {
	C.gtk_notebook_next_page(NOTEBOOK(v))
}

func (v *Notebook) PrevPage() {
	C.gtk_notebook_prev_page(NOTEBOOK(v))
}

func (v *Notebook) ReorderChild(child IWidget, position int) {
	C.gtk_notebook_reorder_child(NOTEBOOK(v), ToNative(child), gint(position))
}

func (v *Notebook) SetTabPos(pos PositionType) {
	C.gtk_notebook_set_tab_pos(NOTEBOOK(v), C.GtkPositionType(pos))
}

func (v *Notebook) SetShowTabs(show_tabs bool) {
	C.gtk_notebook_set_show_tabs(NOTEBOOK(v), gbool(show_tabs))
}

func (v *Notebook) SetShowBorder(show_border bool) {
	C.gtk_notebook_set_show_border(NOTEBOOK(v), gbool(show_border))
}

func (v *Notebook) SetScrollable(scrollable bool) {
	C.gtk_notebook_set_scrollable(NOTEBOOK(v), gbool(scrollable))
}

func (v *Notebook) PopupEnable() {
	C.gtk_notebook_popup_enable(NOTEBOOK(v))
}

func (v *Notebook) PopupDisable() {
	C.gtk_notebook_popup_disable(NOTEBOOK(v))
}

func (v *Notebook) GetCurrentPage() int {
	return int(C.gtk_notebook_get_current_page(NOTEBOOK(v)))
}

func (v *Notebook) GetMenuLabel(child IWidget) *Widget {
	return &Widget{
		C.gtk_notebook_get_menu_label(NOTEBOOK(v), ToNative(child))}
}

func (v *Notebook) GetNthPage(page_num int) *Widget {
	return &Widget{
		C.gtk_notebook_get_nth_page(NOTEBOOK(v), gint(page_num))}
}

func (v *Notebook) GetNPages() int {
	return int(C.gtk_notebook_get_n_pages(NOTEBOOK(v)))
}

func (v *Notebook) GetTabLabel(child IWidget) *Widget {
	return &Widget{
		C.gtk_notebook_get_tab_label(NOTEBOOK(v), ToNative(child))}
}

func (v *Notebook) SetMenuLabel(child IWidget, menu_label IWidget) {
	C.gtk_notebook_set_menu_label(NOTEBOOK(v), ToNative(child), ToNative(menu_label))
}

func (v *Notebook) SetMenuLabelText(child IWidget, name string) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_notebook_set_menu_label_text(NOTEBOOK(v), ToNative(child), gstring(ptr))
}

func (v *Notebook) SetTabLabel(child IWidget, tab_label IWidget) {
	C.gtk_notebook_set_tab_label(NOTEBOOK(v), ToNative(child), ToNative(tab_label))
}

func (v *Notebook) SetTabLabelText(child IWidget, name string) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_notebook_set_tab_label_text(NOTEBOOK(v), ToNative(child), gstring(ptr))
}

func (v *Notebook) SetReorderable(child IWidget, reorderable bool) {
	C.gtk_notebook_set_tab_reorderable(NOTEBOOK(v), ToNative(child), gbool(reorderable))
}

func (v *Notebook) SetTabDetachable(child IWidget, detachable bool) {
	C.gtk_notebook_set_tab_detachable(NOTEBOOK(v), ToNative(child), gbool(detachable))
}

func (v *Notebook) GetMenuLabelText(child IWidget) string {
	return gostring(C.gtk_notebook_get_menu_label_text(NOTEBOOK(v), ToNative(child)))
}

func (v *Notebook) GetScrollable() bool {
	return gobool(C.gtk_notebook_get_scrollable(NOTEBOOK(v)))
}

func (v *Notebook) GetShowBorder() bool {
	return gobool(C.gtk_notebook_get_show_border(NOTEBOOK(v)))
}

func (v *Notebook) GetShowTabs() bool {
	return gobool(C.gtk_notebook_get_show_tabs(NOTEBOOK(v)))
}

func (v *Notebook) GetTabLabelText(child IWidget) string {
	return gostring(C.gtk_notebook_get_tab_label_text(NOTEBOOK(v), ToNative(child)))
}

func (v *Notebook) GetTabPos() uint {
	return uint(C.gtk_notebook_get_tab_pos(NOTEBOOK(v)))
}

func (v *Notebook) GetTabReorderable(child IWidget) bool {
	return gobool(C.gtk_notebook_get_tab_reorderable(NOTEBOOK(v), ToNative(child)))
}

func (v *Notebook) GetTabDetachable(child IWidget) bool {
	return gobool(C.gtk_notebook_get_tab_detachable(NOTEBOOK(v), ToNative(child)))
}

func (v *Notebook) SetCurrentPage(pageNum int) {
	C.gtk_notebook_set_current_page(NOTEBOOK(v), gint(pageNum))
}

func (v *Notebook) SetGroupName(group string) {
	panic_if_version_older(2, 24, 0, "gtk_notebook_set_group_name()")
	ptr := C.CString(group)
	defer cfree(ptr)
	C._gtk_notebook_set_group_name(NOTEBOOK(v), gstring(ptr))
}

func (v *Notebook) GetGroupName() string {
	panic_if_version_older(2, 24, 0, "gtk_notebook_get_group_name()")
	return gostring(C._gtk_notebook_get_group_name(NOTEBOOK(v)))
}
