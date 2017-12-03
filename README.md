# go-gtk3


Fork from:
https://github.com/mattn/go-gtk and upgraded to use gtk3


### Minimal Support of gtk3 and gdk3 on linux!

Priority is to run: `/example/demo/demo.go`

### ToDo
 - [x] Run demo.go in gtk3
 - [ ] Clean `deprecated` warnings
 - [x] Rename `gtk` package to `gtk3`
 - [x] Rename `gdk` package to `gdk3`
 - [ ] Split `gtk3` to separate files (11000 lines in one file is to much)
 - [ ] Split `gdk3` to separate files
 - [ ] Write more tests for widgets


# Implemented:  

## Layout
[Layout Containers](https://developer.gnome.org/gtk3/3.22/LayoutContainers.html)

| Pack/Widget  | Implemented |
| ------------- | ------------- |
| GtkBox            | 9 / 14 |
| GtkGrid           | 0 / 21 |
| GtkRevealer       | 0 / 8 |
| GtkListBox        | 0 / 43 |
| GtkFlowBox        | 0 / 39 |
| GtkStack          | 0 / 22 |
| GtkStackSwitcher  | 0 / 3 |
| GtkStackSidebar   | 0 / 3 |
| GtkActionBar      | 0 / 5 |
| GtkHeaderBar      | 0 / 15 |
| GtkOverlay        | 0 / 5 |
| GtkButtonBox      | 0 / 7 |
| GtkPaned          | 8 / 12 |
| GtkLayout         | 10 / 10 |
| GtkNotebook       | 40 / 45 |
| GtkExpander       | 14 / 18 |
| GtkOrientable     | 0 / 2 |
| GtkAspectFrame    | 0 / 2 |
| GtkFixed          | 3 / 3 |


# Legacy


`go-gtk` code authors:

    Yasuhiro Matsumoto

    David Roundy
    Mark Andrew Gerads
    Tobias Kortkamp
    Mikhail Trushnikov
    Federico Sogaro
    Crazy2be
    Daniël de Kok
    Erik Lissel
    Jeffrey Bolle
    Leonhard Küper
    Matt Joiner
    SQP
    Steven T
    Taru Karttunen
    Utkan Güngördü
    matiaslina
    Dag Robøle
    Denis Dyakov
    Giuseppe Mazzotta
