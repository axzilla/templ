// Code generated by templ DO NOT EDIT.

package example

import "github.com/a-h/templ"
import "context"
import "io"
import "fmt"
import "time"

func headerTemplate(name string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		err = templ.RenderScripts(ctx, w, )
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<header")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " data-testid=\"headerTemplate\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<h1>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(name))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</h1>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</header>")
		if err != nil {
			return err
		}
		return err
	})
}

func footerTemplate() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		err = templ.RenderScripts(ctx, w, )
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<footer")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " data-testid=\"footerTemplate\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div>")
		if err != nil {
			return err
		}
		var_1 := `&copy; `
		_, err = io.WriteString(w, var_1)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(fmt.Sprintf("%d", time.Now().Year())))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</footer>")
		if err != nil {
			return err
		}
		return err
	})
}

