package rectangles
import "fmt"
// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawRectangle(height, width int, inner, outer string) {
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if h == 0 || h == height-1 || w == 0 || w == width-1 {
				fmt.Print(outer)
			} else {
				fmt.Print(inner)
			}
		}
		fmt.Println()
	}
}

// REMARKS
// - Wenn Sie diese Aufgabe gelöst haben, können Sie die Aufgaben
//   für das Zeichnen von leeren und gefüllten Rechtecken sehr viel einfacher lösen.
