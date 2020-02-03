package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (a [][]uint8) {
    a = make([][]uint8, dy)

    for i := range a {
        a[i] = make([]uint8, dx)
    }

    for i := range a {
        for j := range a[i] {
            switch {
            case j % 15 == 0:
                a[i][j] = 240
            case j % 3 == 0:
                a[i][j] = 120
            case j % 5 == 0:
                a[i][j] = 150
            default:
                a[i][j] = 100
            }
        }
    }

    return
}

func main() {
	pic.Show(Pic)
}

