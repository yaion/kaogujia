package v1

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/gzip"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{gzip.Gzip(gzip.DefaultCompression)}
}

func _v1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deleteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deleteuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateuserMw() []app.HandlerFunc {
	// your code...
	return nil
}
