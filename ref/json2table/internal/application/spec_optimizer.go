/*
Package application run the application
*/
package application

func optimizeSpec(spec *Spec) {
	for i, column := range spec.Columns {
		if column.MinWidth > 0 && column.Width < column.MinWidth {
			spec.Columns[i].Width = column.MinWidth
		}
		if column.MaxWidth > 0 && column.Width > column.MaxWidth {
			spec.Columns[i].Width = column.MaxWidth
		}
	}
}
